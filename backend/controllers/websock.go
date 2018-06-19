package controllers

import (
	"github.com/astaxie/beego"
	"go-chat/backend/models"
)

// Operations about object
type WebsockController struct {
	beego.Controller
}

func (o *WebsockController) Get() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}
