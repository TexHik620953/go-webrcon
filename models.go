package webrcon

type ServerInfo struct {
	Hostname          string  `json:"Hostname"`
	MaxPlayers        int     `json:"MaxPlayers"`
	Players           int     `json:"Players"`
	Queued            int     `json:"Queued"`
	Joining           int     `json:"Joining"`
	ReservedSlots     int     `json:"ReservedSlots"`
	EntityCount       int     `json:"EntityCount"`
	GameTime          string  `json:"GameTime"`
	Uptime            int     `json:"Uptime"`
	Map               string  `json:"Map"`
	Framerate         float64 `json:"Framerate"`
	Memory            int     `json:"Memory"`
	MemoryUsageSystem int     `json:"MemoryUsageSystem"`
	Collections       int     `json:"Collections"`
	NetworkIn         int     `json:"NetworkIn"`
	NetworkOut        int     `json:"NetworkOut"`
	Restarting        bool    `json:"Restarting"`
	SaveCreatedTime   string  `json:"SaveCreatedTime"`
	Version           int     `json:"Version"`
	Protocol          string  `json:"Protocol"`
}

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
	Type       int    `json:"Type"`
}

type Report struct {
	*Feedback
	TargetId   string `json:"TargetId"`
	TargetName string `json:"TargetName"`
}
