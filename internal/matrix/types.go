package matrix

type Credentials struct {
	Username string
	Password string
}

type LoginRequestData struct {
	Type       string   `json:"type"`
	Identifier UserData `json:"identifier"`
	Password   string   `json:"password"`
}

type LoginResponseData struct {
	AccessToken  string `json:"access_token"`
	DeviceID     string `json:"device_id"`
	ExpiresInMs  int    `json:"expires_in_ms"`
	RefreshToken string `json:"refresh_token"`
	UserID       string `json:"user_id"`
}

type UserData struct {
	Type string `json:"type"`
	User string `json:"user"`
}

type MessageData struct {
	MsgType string `json:"msgtype"`
	Body    string `json:"body"`
}
