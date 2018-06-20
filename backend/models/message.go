package models

// 接受消息结构
type ReceiveMessage struct {
	Action   string    `json:"action"`
	Token    string    `json:"token"`
	Message    string    `json:"message"`
}

// 发送消息结构
type SendMessage struct {
	Action   string    `json:"action"`
	Data MessageDetail `json:"data"`
}

type MessageDetail struct {
	Token    string    `json:"token"`
	Message    string    `json:"message"`
}


