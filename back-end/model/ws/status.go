package ws

type Status struct {
	UUID   string `json:"uuid"`
	Target string `json:"target"`
	Value  any    `json:"value"`
}
