package ws

type Msg struct {
	Content string `json:"content"`
	Type    string `json:"type"`
	From    string `json:"from"`
	ToType  string `json:"toType"`
	To      string `json:"to"`
}

type Response struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}
