package controllers

import (
	"fyoukuApi/models"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) ChannelRegion() {
	channelId, _ := c.GetInt("channelId")
	if channelId == 0 {
		c.Data["json"] = ReturnError(4001,"必须指定频道")
	}
	num,regions,err := models.GetChannelRegion(channelId)
	if err == nil{
		c.Data["json"] = ReturnSuccess(0,"success",regions,num)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(5000,"没有相关内容")
		c.ServeJSON()
	}
}

func (c *BaseController) ChannelType() {
	channelId, _ := c.GetInt("channelId")
	if channelId == 0 {
		c.Data["json"] = ReturnError(4001,"必须指定频道")
	}
	num,types,err := models.GetChannelType(channelId)
	if err == nil{
		c.Data["json"] = ReturnSuccess(0,"success",types,num)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(5000,"没有相关内容")
		c.ServeJSON()
	}
}