package models

import (
	"log"
	"gitHub.com/gorilla/websocket"
	"encoding/json"
	"gitHub.com/astaxie/beego"
)

const (
	// Maximum message size allowed from peer.
	MaxMessageSize = 512
)


// Client is a middleman between the websocket Connection and the Hub.
type Client struct {
	Hub *Hub

	// The websocket Connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte
}

// readPump pumps messages from the websocket Connection to the Hub.
//
// The application runs readPump in a per-Connection goroutine. The application
// ensures that there is at most one reader on a Connection by executing all
// reads from this goroutine.
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.UnRegister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(MaxMessageSize)
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		c.Hub.Broadcast <- message
	}
}

// writePump pumps messages from the Hub to the websocket Connection.
//
// A goroutine running writePump is started for each Connection. The
// application ensures that there is at most one writer to a Connection by
// executing all writes from this goroutine.
func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				beego.Error("发生异常，关闭链接",c)
				// The Hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			token := GetRandomString(10)
			receiveMessage   := &ReceiveMessage{}
			sendMessage   := &SendMessage{}

			// 解析读
			if err := json.Unmarshal(message, receiveMessage); err == nil {
				switch receiveMessage.Action {
				case "open":
					// 需要返回数据
					sendMessage.Data.Token = token
					sendMessage.Action = receiveMessage.Action
					// 返回数据json 为二进制
					if 	returnData, err := json.Marshal(sendMessage) ; err == nil {
						_, err = w.Write(returnData)
						if err != nil {
							beego.Error("write:", err)
							break
						} else {
							beego.Info("建立链接：" + token," data:",sendMessage)
						}
					}

				case "sendMessage":
					// 需要返回数据
					sendMessage.Action = "replyMessage"
					sendMessage.Data.Token = receiveMessage.Token
					sendMessage.Data.Message = receiveMessage.Message
					// 返回数据json 为二进制
					if 	returnData, err := json.Marshal(sendMessage) ; err == nil {
						_, err = w.Write(returnData)
						if err != nil {
							beego.Info("send " + token + ": error", err)
							break
						} else {
							beego.Info("send " + token + ": success", sendMessage)
						}
					}
				// 关闭
				case "close":
					c.Hub.UnRegister <- c
				}
			}

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}
