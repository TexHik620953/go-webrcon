package webrcon

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/TexHik620953/go-webrcon/events"
	"github.com/TexHik620953/go-webrcon/utils"

	"github.com/gorilla/websocket"
)

var ErrTimeout = errors.New("rcon response timeout reached")

type WebRconClient struct {
	ctx    context.Context
	dialer *websocket.Dialer
	conn   *websocket.Conn

	lastId          int64
	responseMap     map[int64]chan *Message
	responseMapSync sync.Mutex

	messageHandlers  *events.EventHandlersGroup[*Message]
	feedbackHandlers *events.EventHandlersGroup[*Feedback]
	reportHandlers   *events.EventHandlersGroup[*Report]
}

func Connect(ctx context.Context, addr string, password string) (*WebRconClient, error) {
	dialer := &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
	}

	wsUrl := fmt.Sprintf("ws://%s/%s", addr, password)
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, wsUrl, nil)
	if err != nil {
		return nil, err
	}

	h := &WebRconClient{
		dialer:           dialer,
		conn:             conn,
		ctx:              ctx,
		lastId:           0,
		responseMap:      map[int64]chan *Message{},
		messageHandlers:  events.NewEventHandlersGroup[*Message](),
		feedbackHandlers: events.NewEventHandlersGroup[*Feedback](),
		reportHandlers:   events.NewEventHandlersGroup[*Report](),
	}

	go h.listenWorker()

	return h, nil
}

func (h *WebRconClient) Exec(msg string, execTimeout time.Duration) (*Message, error) {
	h.lastId++
	packet := &CommandPacket{
		Message: msg,
		Name:    "WebRcon",
	}

	ch := make(chan *Message, 1)

	h.responseMapSync.Lock()
	h.lastId++
	h.responseMap[h.lastId] = ch
	packet.Identifier = h.lastId
	h.responseMapSync.Unlock()

	defer func() {
		h.responseMapSync.Lock()
		delete(h.responseMap, h.lastId)
		h.responseMapSync.Unlock()
	}()

	err := h.conn.WriteJSON(packet)
	if err != nil {
		return nil, err
	}

	select {
	case <-time.After(execTimeout):
		return nil, ErrTimeout
	case result := <-ch:
		return result, nil
	}
}

func (h *WebRconClient) listenWorker() {
	var err error
	tempMsg := &messageInternal{}
	for {
		err = h.conn.ReadJSON(tempMsg)
		if err != nil {
			if utils.IsWsTimeout(err) {
				continue
			}
			log.Printf("failed to unmarshal msg: %s", err.Error())
			continue
		}

		msg, err := tempMsg.Normalize()
		if err != nil {
			log.Printf("failed to normalize msg: %s", err.Error())
			continue
		}
		h.responseMapSync.Lock()
		ch, ex := h.responseMap[msg.Identifier]
		h.responseMapSync.Unlock()
		// Callback
		if ex {
			ch <- msg
			continue
		}

		// Report
		if msg.Type == MESSAGE_TYPE_REPORT {
			rep := &Report{}
			err = json.Unmarshal([]byte(msg.Message), rep)
			if err != nil {
				log.Printf("failed to unmarshal msg report: %s", err.Error())
				continue
			}
			// Feedback if no target
			if len(rep.TargetId) == 0 {
				h.emitFeedback(rep.Feedback)
			} else {
				h.emitReport(rep)
			}
			continue
		}
		// Regular message
		if msg.Identifier == 0 {
			h.emitMessage(msg)
		}
	}
}
