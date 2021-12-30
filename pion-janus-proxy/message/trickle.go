package message

type TrickleRequest struct {
	Request   string      `json:"janus" mapstructure:"janus"`
	Session   int         `json:"session_id" mapstructure:"session_id"`
	Handle    int         `json:"handle_id" mapstructure:"handle_id"`
	Candidate interface{} `json:"candidate" mapstructure:"candidate"`
}

type TrickleRequestComplete struct {
	Completed bool `json:"completed" mapstructure:"completed"`
}
