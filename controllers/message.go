package controllers

import (
	"beego-video/models"
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
			close(sendChan)
		}()

		for i :=0; i<5; i++ {
			go sendMessageFunc(sendChan,closeChan)
		}

		for i := 0; i < 5; i++ {
			<-closeChan
		}
		close(closeChan)

		c.Data["json"] = ReturnSuccess(0,"发送成功~","",1)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(5000, "发送失败，请联系客服~")
		c.ServeJSON()
	}
}

func sendMessageFunc(sendChan chan SendData,closeChan chan bool)  {
	for t := range sendChan {
		models.SendMessageUserMq(t.UserId, t.MessageId)
	}
	closeChan <- true
}