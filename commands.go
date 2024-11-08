package webrcon

import (
	"encoding/json"
	"time"
)

func (h *WebRconClient) ListPlayers(execTimeout time.Duration) ([]*RconPlayer, error) {
	resp, err := h.Exec("playerlist", execTimeout)
	if err != nil {
		return nil, err
	}
	data := make([]*RconPlayer, 0)
	err = json.Unmarshal([]byte(resp.Message), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (h *WebRconClient) ServerInfo(execTimeout time.Duration) (*ServerInfo, error) {
	resp, err := h.Exec("serverinfo", execTimeout)
	if err != nil {
		return nil, err
	}
	data := &ServerInfo{}
	err = json.Unmarshal([]byte(resp.Message), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
