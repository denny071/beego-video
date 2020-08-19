package controllers

import (
	"fyoukuApi/models"
	"github.com/astaxie/beego"
)


type Comment struct {
	Id           int             `json:"id"`
	Content      string          `json:"content"`
	AddTime      int64           `json:"addTime"`
	AddTimeTitle string          `json:"addTimeTitle"`
	UserId       int             `json:"userId"`
	Stamp        int             `json:"stamp"`
	PraiseCount  int             `json:"praiseCount"`
	UserInfo     models.UserInfo `json:"userinfo"`
	EpisodesId   int             `json:"episodesId"`
}

type CommentController struct {
	beego.Controller
}

func (c *CommentController) List() {
	c.Data["json"] = SuccessJsonStruct{Msg: "List"}
	c.ServeJSON()
}

func (c *CommentController) Save() {
	c.Data["json"] = SuccessJsonStruct{Msg: "Save"}
	c.ServeJSON()
}
