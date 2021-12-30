package message

type SetupResponse struct {
	Janus       string     `json:"janus"`
	SessionID   int64      `json:"session_id"`
	Transaction string     `json:"transaction"`
	Sender      int64      `json:"sender"`
	Plugindata  Plugindata `json:"plugindata"`
	Jsep        Jsep       `json:"jsep"`
}

type Data struct {
	Textroom string `json:"textroom"`
	Result   string `json:"result"`
}
type Plugindata struct {
	Plugin string `json:"plugin"`
	Data   Data   `json:"data"`
}
type Jsep struct {
	Type string `json:"type"`
	Sdp  string `json:"sdp"`
}
