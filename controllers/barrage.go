package controllers

import "github.com/astaxie/beego"

type BarrageController struct {
	beego.Controller
}

func (c *BarrageController) List() {
	c.Data["json"] = SuccessJsonStruct{Msg: "get Barrage"}
	c.ServeJSON()
}

func (c *BarrageController) Save() {
	c.Data["json"] = SuccessJsonStruct{Msg: "Save Barrage"}
	c.ServeJSON()
}