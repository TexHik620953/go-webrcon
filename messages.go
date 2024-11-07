package webrcon

type messageInternal struct {
	Message    string `json:"Message"`
	Identifier int64  `json:"Identifier"`
	Type       string `json:"Type"`
	Stacktrace string `json:"Stacktrace"`
}

// Generic  Warning
func (h *messageInternal) Normalize() (*Message, error) {
	r := &Message{
		Message:    h.Message,
		Stacktrace: h.Stacktrace,
		Identifier: h.Identifier,
		Type:       MessageType(h.Type),
	}

	return r, nil
}

type MessageType string

const (
	MESSAGE_TYPE_GENERIC    MessageType = "Generic"
	MESSAGE_TYPE_ERROR      MessageType = "Error"
	MESSAGE_TYPE_WARNING    MessageType = "Warning"
	MESSAGE_TYPE_CHAT       MessageType = "Chat"
	MESSAGE_TYPE_REPORT     MessageType = "Report"
	MESSAGE_TYPE_CLIENTPERF MessageType = "ClientPerf"
)

type Message struct {
	Message    string      `json:"Message"`
	Identifier int64       `json:"Identifier"`
	Type       MessageType `json:"Type"`
	Stacktrace string      `json:"Stacktrace"`
}

type CommandPacket struct {
	Message    string `json:"Message"`
	Identifier int64  `json:"Identifier"`
	Name       string `json:"Name"`
}
