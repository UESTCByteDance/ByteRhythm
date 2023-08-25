package mq

import (
	"ByteRhythm/config"
	"fmt"

	"github.com/streadway/amqp"
)

var RabbitMQ *amqp.Connection

func InitRabbitMQ() {
	config.Init()
	connString := fmt.Sprintf("%s://%s:%s@%s:%s/",
		config.RabbitMQ,
		config.RabbitMQUser,
		config.RabbitMQPassWord,
		config.RabbitMQHost,
		config.RabbitMQPort,
	)
	conn, err := amqp.Dial(connString)
	if err != nil {
		panic(err)
	}
	RabbitMQ = conn
}
