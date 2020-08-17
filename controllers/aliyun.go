package controllers

import "github.com/astaxie/beego"

type AliyunController struct {
	beego.Controller
}

func (c * AliyunController) UploadVideo()  {
	c.Data["json"] = SuccessJsonStruct{Msg: "UploadVideo"}
	c.ServeJSON()
}

func (c * AliyunController) PlayAuth()  {
	c.Data["json"] = SuccessJsonStruct{Msg: "PlayAuth"}
	c.ServeJSON()
}

func (c * AliyunController) Callback()  {
	c.Data["json"] = SuccessJsonStruct{Msg: "callback"}
	c.ServeJSON()
}