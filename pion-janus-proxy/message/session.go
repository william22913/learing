package message

type SessionAttachRequest struct {
	Request string `json:"janus" mapstructure:"janus"`
	Session int    `json:"session_id" mapstructure:"session_id"`
	Plugin  string `json:"plugin" mapstructure:"plugin"`
	Tag     string `json:"tag" mapstructure:"tag"`
}

type SessionData struct {
	ID int `json:"id" mapstructure:"id"`
}
