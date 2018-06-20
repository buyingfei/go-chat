package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"go-chat/backend/models"
	"net/http"
	"strings"
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

	ws.SetCloseHandler(func(code int, text string) error {
		ip := user[ws]
		beego.Info("出现错误,关闭链接：" + ip)
		deleteClinet(ws)
		return nil
	})

	// 获取ip
	ip:=this.Ctx.Request.RemoteAddr
	ip=ip[0:strings.LastIndex(ip, ":")]

	user[ws] = ip
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			beego.Info("read ip" + ip +" error:", err)
			ws.Close()
			deleteClinet(ws)
			break
		}
		messStruct   := &models.MessageStruct{}
		if err := json.Unmarshal(message, messStruct); err == nil {
			switch messStruct.Action {
			case "open":
				// 发送数据给前端,将ip 作为token
				messStruct.Data.Token = "ip" + ip
				if 	returnData, err := json.Marshal(messStruct) ; err == nil {
					err = ws.WriteMessage(mt, returnData)
					if err != nil {
						beego.Error("write:", err)
						break
					}
					beego.Info("建立链接：" + "ip" + ip)
				}
			case "sendMessage":
				// 将message 发送到各个链接
				for nowws,token := range user {
					if token == messStruct.Data.Token {
						continue
					}
					messStruct.Action = "replyMessage"
					messStruct.Data.Token = messStruct.Token
					messStruct.Data.Message = messStruct.Message
					if 	returnData, err := json.Marshal(messStruct) ; err == nil {
						err = nowws.WriteMessage(mt, returnData)
						if err != nil {
							beego.Info("send " + token + ": error", err)
							break
						}
					}
				}
			case "close":
				deleteClinet(ws)
			}

		}

	}
}
