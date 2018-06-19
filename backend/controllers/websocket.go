package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	"net/http"
)

// Operations about object
type WebsocketController struct {
	beego.Controller
}


// Join method handles WebSocket requests for WebSocketController.
func (this *WebsocketController) Get() {

	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			beego.Error("read:", err)
			break
		}
		err = ws.WriteMessage(mt, message)
		if err != nil {
			logs.Info("write:", err)
			break
		}
	}
}
