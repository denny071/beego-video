package mq

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
)

type Callback func(msg string)

func Connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	return conn, err
}

func Publish(exchange string, queueName string, body string) error {

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
	err = channel.Publish(exchange, q.Name, false, false, amqp.Publishing{
		DeliveryMode:    amqp.Persistent,
		ContentEncoding: "text/plain",
		Body:            []byte(body),
	})

	return err
}

func Consumer(exchange string, queueName string, callback Callback) {
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
	msgs, err := channel.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			s := BytesToString(&(d.Body))
			callback(*s)
			d.Ack(false)
		}
	}()
	<-forever
}

func PublishDlx(exchangeA string, body string) error {
	//建立连接
	conn, err := Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	//创建一个Channel
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	//消息发送到A交换机
	err = channel.Publish(exchangeA, "", false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(body),
	})

	return err
}
func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}
