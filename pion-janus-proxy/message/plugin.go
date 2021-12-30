package message

type PluginData struct {
	Data map[string]interface{} `json:"data" mapstructure:"data"`
}

type PluginResponse struct {
	Status      string     `json:"janus" mapstructure:"janus"`
	Session     int        `json:"session_id" mapstructure:"session_id"`
	Handle      int        `json:"sender" mapstructure:"sender"`
	Transaction string     `json:"transaction" mapstructure:"transaction"`
	Data        PluginData `json:"plugindata" mapstructure:"plugindata"`
}

type PluginRequest struct {
	Request string      `json:"janus" mapstructure:"janus"`
	Session int         `json:"session_id" mapstructure:"session_id"`
	Handle  int         `json:"handle_id" mapstructure:"handle_id"`
	Body    interface{} `json:"body" mapstructure:"body"`
}

type PluginAckRequest struct {
	Request string      `json:"janus" mapstructure:"janus"`
	Session int         `json:"session_id" mapstructure:"session_id"`
	Handle  int         `json:"handle_id" mapstructure:"handle_id"`
	Body    interface{} `json:"body" mapstructure:"body"`
	JSEP    JSEP        `json:"jsep" mapstructure:"jsep"`
}

type JSEP struct {
	Type string `json:"type" mapstructure:"type"`
	SDP  string `json:"sdp" mapstructure:"sdp"`
}

type PluginHandle struct {
	ID int `json:"id" mapstructure:"id"`
}
