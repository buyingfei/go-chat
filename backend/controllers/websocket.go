package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
)

// Operations about object
type WebsocketController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{
} // use default options

// Join method handles WebSocket requests for WebSocketController.
func (this *WebsocketController) Get() {
	// 允许跨域
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	// // 建立socket 链接
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			beego.Error("read:", err)
			break
		}
		err = ws.WriteMessage(mt, message)
		if err != nil {
			beego.Info("write:", err)
			break
		}
	}
}
