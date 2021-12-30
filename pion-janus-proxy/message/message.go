package message

type JanusRequest struct {
	Request     string `json:"janus" mapstructure:"janus"`
	Transaction string `json:"transaction" mapstructure:"transaction"`
}

type JanusResponse struct {
	Status      string                 `json:"janus"`
	Transaction string                 `json:"transaction"`
	Data        map[string]interface{} `json:"data"`
	Message     string                 `json:"message"`
}

type HandshakeRequest struct {
	UserId      string `json:"user_id" mapstructure:"user_id"`
	Role        string `json:"role" mapstructure:"role"`
	Conference  string `json:"conference_id" mapstructure:"conference_id"`
	Transaction string `json:"transaction" mapstructure:"transaction"`
	Token       string `json:"token" mapstructure:"token"`
	Tag         string `json:"tag" mapstructure:"tag"`
}

func IsAuthorized(resp JanusResponse) bool {
	return resp.Status == "authorized"
}
