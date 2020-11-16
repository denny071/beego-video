package models

import (
	"encoding/json"
	mq "beego-video/services/mp"
	"github.com/astaxie/beego/orm"
	"time"
)

type Message struct {
	Id int
	Content string
	AddTime int64
}

func init() {
	orm.RegisterModel(new(Message))
}

func SendMessageDo(content string) (int64, error)  {
	o := orm.NewOrm()
	var message Message
	message.Content = content
	message.AddTime = time.Now().Unix()
	return o.Insert(&message)
}

func SendMessageUserMq(userId int, messageId int64)  {
	type Data struct {
		UserId int
		MessageId int64
	}
	var data Data
	data.UserId = userId
	data.MessageId = messageId
	dataJson, _ := json.Marshal(data)
	mq.Publish("","fyouku_send_message_user",string(dataJson))
}