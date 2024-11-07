package webrcon

type RconPlayer struct {
	SteamID          string  `json:"SteamID"`
	OwnerSteamID     string  `json:"OwnerSteamID"`
	DisplayName      string  `json:"DisplayName"`
	Ping             int     `json:"Ping"`
	Address          string  `json:"Address"`
	ConnectedSeconds float64 `json:"ConnectedSeconds"`
	VoiationLevel    float64 `json:"VoiationLevel"`
	CurrentLevel     float64 `json:"CurrentLevel"`
	UnspentXp        float64 `json:"UnspentXp"`
	Health           float64 `json:"Health"`
}

type Feedback struct {
	PlayerId   string `json:"PlayerId"`
	PlayerName string `json:"PlayerName"`
	Subject    string `json:"Subject"`
	Message    string `json:"Message"`
	Type       string `json:"Type"`
}

type Report struct {
	*Feedback
	TargetId   string `json:"TargetId"`
	TargetName string `json:"TargetName"`
}
