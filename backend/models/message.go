package models

type MessageStruct struct {
	Action   string    `json:"action"`
	// send
	Data MessageDetail `json:"data"`
	// receive
	Token    string    `json:"token"`
	Message    string    `json:"message"`
}

type MessageDetail struct {
	Token    string    `json:"token"`
	Message    string    `json:"message"`
}


