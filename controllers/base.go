package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (c *BaseController) ChannelRegion() {
	c.Data["json"] = SuccessJsonStruct{Msg: "get ChannelRegion"}
	c.ServeJSON()
}

func (c *BaseController) ChannelType() {
	c.Data["json"] = SuccessJsonStruct{Msg: "get ChannelType"}
	c.ServeJSON()
}