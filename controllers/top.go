package controllers

import "github.com/astaxie/beego"

type TopController struct {
	beego.Controller
}

func (c *TopController) ChannelTop() {
	c.Data["json"] = SuccessJsonStruct{Msg: "ChannelTop"}
	c.ServeJSON()
}

func (c *TopController) TypeTop() {
	c.Data["json"] = SuccessJsonStruct{Msg: "TypeTop"}
	c.ServeJSON()
}