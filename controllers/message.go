package controllers

import "github.com/astaxie/beego"

type MessageController struct {
	beego.Controller
}

func (c * MessageController) Send()  {
	c.Data["json"] = SuccessJsonStruct{Msg: "send"}
	c.ServeJSON()
}