package message

// {"status":false,"message":"INVALID_SESSION_ID"}

type JanusMessageError struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
