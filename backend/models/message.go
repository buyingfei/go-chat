package models

type MessageStruct struct {
	Action   string    `json:"action"`
	Data MessageDetail `json:"data"`
}

type MessageDetail struct {
	Token    string    `json:"token"`
	Message    string    `json:"message"`
}


