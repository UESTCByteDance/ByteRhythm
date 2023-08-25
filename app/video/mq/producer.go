package mq

import (
	"github.com/streadway/amqp"
)

func SendMessage2MQ(body []byte, queueName string) (err error) {
	ch, err := RabbitMQ.Channel()
	if err != nil {
		return
	}
	q, _ := ch.QueueDeclare(queueName, true, false, false, false, nil)
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		return
	}
	return
}
