package mq

import (
	"github.com/streadway/amqp"
)

func CreateFavorite2MQ(body []byte) (err error) {
	ch, err := RabbitMQ.Channel()
	if err != nil {
		return
	}
	q, _ := ch.QueueDeclare("favorite-create-queue", true, false, false, false, nil)
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

func DeleteFavorite2MQ(body []byte) (err error) {
	ch, err := RabbitMQ.Channel()
	if err != nil {
		return
	}
	q, _ := ch.QueueDeclare("favorite-delete-queue", true, false, false, false, nil)
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
