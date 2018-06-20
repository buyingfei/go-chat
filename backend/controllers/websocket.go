package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"go-chat/backend/models"
	"net/http"
	"encoding/json"
	"sync"
)

// Operations about object
type WebsocketController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{
	// 允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

// 定义用户
var userLock sync.Mutex
var user = make(map[*websocket.Conn] string)

func deleteClinet(ws *websocket.Conn) {
	userLock.Lock()
	defer userLock.Unlock()
	if _, ok := user[ws]; ok {
		beego.Info("关闭链接：",user[ws])
		delete(user, ws)
	}
}


// Join method handles WebSocket requests for WebSocketController.
func (this *WebsocketController) Get() {

	// 建立socket 链接
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	defer ws.Close()

	// 关闭链接处理
	ws.SetCloseHandler(func(code int, text string) error {
		deleteClinet(ws)
		return nil
	})

	token := models.GetRandomString(8)
	userLock.Lock()
	defer userLock.Unlock()
	user[ws] = token

	for {
		// 读消息
		mt, message, err := ws.ReadMessage()
		if err != nil {
			// 读失败关闭链接，删除map对象
			beego.Info("read " + token +" error:", err)
			deleteClinet(ws)
			ws.Close()
			break
		}

		receiveMessage   := &models.ReceiveMessage{}
		sendMessage   := &models.SendMessage{}

		// 解析读
		if err := json.Unmarshal(message, receiveMessage); err == nil {
			switch receiveMessage.Action {
			case "open":
				// 需要返回数据
				sendMessage.Data.Token = token
				sendMessage.Action = receiveMessage.Action
				// 返回数据json 为二进制
				if 	returnData, err := json.Marshal(sendMessage) ; err == nil {
					err = ws.WriteMessage(mt, returnData)
					if err != nil {
						beego.Error("write:", err)
						break
					} else {
						beego.Info("建立链接：" + token," data:",sendMessage)
					}
				}

			case "sendMessage":
				// 将message 发送到各个链接
				for nowws,token := range user {
					if token == receiveMessage.Token {
						continue
					}
					beego.Info(token,receiveMessage.Token)
					// 需要返回数据
					sendMessage.Action = "replyMessage"
					sendMessage.Data.Token = receiveMessage.Token
					sendMessage.Data.Message = receiveMessage.Message
					// 返回数据json 为二进制
					if 	returnData, err := json.Marshal(sendMessage) ; err == nil {
						err = nowws.WriteMessage(mt, returnData)
						if err != nil {
							beego.Info("send " + token + ": error", err)
							break
						} else {
							beego.Info("send " + token + ": success", sendMessage)
						}
					}
				}

			case "close":
				deleteClinet(ws)
			}

		}

	}
}
