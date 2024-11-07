package webrcon

import "context"

const (
	MSGEVENT_MESSAGE = "msg"
	MSGEVENT_CHAT    = "chat"
	MSGEVENT_WARNING = "wrn"
	MSGEVENT_ERROR   = "err"

	MSGEVENT_ALL = "all"

	EVENT_REPORT = "rep"
)

func (h *WebRconClient) OnMessage(handler func(*Message)) context.CancelFunc {
	return h.messageHandlers.Add(handler)
}
func (h *WebRconClient) emitMessage(val *Message) {
	h.messageHandlers.Emit(val)
}

func (h *WebRconClient) OnFeedback(handler func(*Feedback)) context.CancelFunc {
	return h.feedbackHandlers.Add(handler)
}
func (h *WebRconClient) emitFeedback(val *Feedback) {
	h.feedbackHandlers.Emit(val)
}

func (h *WebRconClient) OnReport(handler func(*Report)) context.CancelFunc {
	return h.reportHandlers.Add(handler)
}
func (h *WebRconClient) emitReport(val *Report) {
	h.reportHandlers.Emit(val)
}
