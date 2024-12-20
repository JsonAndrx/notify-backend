package types

type Success struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
	Message  interface{} `json:"message"`
}

type Error struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
}