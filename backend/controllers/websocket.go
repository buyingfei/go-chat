package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"go-chat/backend/models"
	"net/http"
	"strings"
	"encoding/json"
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
var user = make(map[*websocket.Conn] string)


// Join method handles WebSocket requests for WebSocketController.
func (this *WebsocketController) Get() {

	// 建立socket 链接
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	defer ws.Close()

	// 获取ip
	ip:=this.Ctx.Request.RemoteAddr
	ip=ip[0:strings.LastIndex(ip, ":")]

	user[ws] = ip

	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			beego.Info("read:", err)
			break
		}
		var messStruct   models.MessageStruct
		if err := json.Unmarshal(message, &messStruct); err == nil {
			beego.Info(messStruct.Action)
			switch messStruct.Action {
			case "open":
				messStruct.Data.Token = ip
			case "sendMessage":
			case "close":
			}
			// 发送数据给前端
		    if 	returnData, err := json.Marshal(messStruct) ; err == nil {
				err = ws.WriteMessage(mt, returnData)
				if err != nil {
					beego.Info("write:", err)
					break
				}
			}
		}

	}
}
