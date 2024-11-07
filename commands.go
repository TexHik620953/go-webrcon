package webrcon

import "encoding/json"

func (h *WebRconClient) ListPlayers() ([]*RconPlayer, error) {
	resp, err := h.Exec("playerlist")
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
