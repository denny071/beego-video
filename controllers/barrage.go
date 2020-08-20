package controllers

import (
	"encoding/json"
	"fyoukuApi/models"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
)

type BarrageController struct {
	beego.Controller
}

type WsData struct {
	CurrentTime int
	EpisodesId int
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func (c *BarrageController) List() {
	var (
		conn *websocket.Conn
		err error
		data []byte
		barrages []models.BarrageData
	)
	if conn, err = upgrader.Upgrade(c.Ctx.ResponseWriter,c.Ctx.Request, nil); err != nil {
		goto ERR
	}
	for {
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		var wsData WsData
		json.Unmarshal([]byte(data), &wsData)
		endTime := wsData.CurrentTime + 60
		// 获取弹幕数据
		_, barrages, err  = models.BarrageList(wsData.EpisodesId, wsData.CurrentTime, endTime)
		if err == nil {
			if err := conn.WriteJSON(barrages); err != nil {
				goto ERR
			}
		}
	}

ERR:
	conn.Close()

}


func (c *BarrageController) Save() {

	episodesId, _ := c.GetInt("episodesId")
	videoId, _ := c.GetInt("videoId")
	userId, _ := c.GetInt("userId")
	currentTime, _ := c.GetInt("currentTime")
	content  := c.GetString("content")
	if videoId == 0 {
		c.Data["json"] = ReturnError(4001,"必须传入视频ID")
		c.ServeJSON()
	}
	if episodesId == 0 {
		c.Data["json"] = ReturnError(4002,"必须传入剧集ID")
		c.ServeJSON()
	}
	if userId == 0 {
		c.Data["json"] = ReturnError(4003,"必须登录")
		c.ServeJSON()
	}
	if currentTime == 0 {
		c.Data["json"] = ReturnError(4004,"必须传入时间")
		c.ServeJSON()
	}
	if content == "" {
		c.Data["json"] = ReturnError(4005,"内容不能为空")
		c.ServeJSON()
	}
	err := models.SaveBarrage(episodesId,videoId,currentTime,userId,content)
	if err == nil{
		c.Data["json"] = ReturnSuccess(0,"success","",1)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(5000,"没有相关内容")
		c.ServeJSON()
	}
}