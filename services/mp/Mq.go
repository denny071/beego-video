package mp

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Callback func(msg string)

func Connect() (*amqp.Connection, error)  {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	return conn, err
}

func createChannel() (*amqp.Channel,error)  {
	// 建立连接
	conn, err := Connect()
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	// 创建通过channel
	channel, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer channel.Close()

	return channel,err
}

func Publish(exchange string, queueName string, body string) error  {

	channel,err := createChannel()
	if err != nil {
		return err
	}
	// 创建队列
	q, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
		)
	if err != nil {
		return err
	}

	// 发送消息
	err  = channel.Publish(exchange,q.Name, false, false,amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentEncoding: "text/plain",
		Body: []byte(body),
	})

	return err
}

func Consumer(exchange string,queueName string, callback Callback)  {
	channel,err := createChannel()
	if err != nil {
		return
	}
	q, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	msgs,err := channel.Consume(q.Name,"", false,false,false,false,nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	forever := make


}