package controllers

import (
	"fyoukuApi/models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type MessageController struct {
	beego.Controller
}
type SendData struct {
	UserId int
	MessageId int64
}

func (c * MessageController) Send()  {
	uids := c.GetString("uids")
	content := c.GetString("content")

	if uids == "" {
		c.Data["json"] = ReturnError(4001,"请填写接收人")
		c.ServeJSON()
	}
	if content == "" {
		c.Data["json"] = ReturnError(4002,"请填写发送内容")
		c.ServeJSON()
	}
	messageId, err := models.SendMessageDo(content)
	if err == nil {
		uidConfig := strings.Split(uids,",")
		count := len(uidConfig)

		sendChan := make(chan SendData, count)
		closeChan := make(chan bool, count)

		go func() {
			var data SendData
			for _,v := range uidConfig {
				userId, _ := strconv.Atoi(v)
				data.UserId = userId
				data.MessageId = messageId
				sendChan <- data
			}
		}()

		for i :=0; i<5; i++ {

		}

	}
}