package mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	QUEUE_PROCESSING = "queue-processing"
	QUEUE_RESPONSE   = "queue-response"
)

func NewRabbitMQ() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	return conn, err
}

func NewChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	channel, err := conn.Channel()
	return channel, err
}

func CloseConnection(conn *amqp.Connection) {
	if err := conn.Close(); err != nil {
		panic(err)
	}
}

func CloseChannel(channel *amqp.Channel) {
	if err := channel.Close(); err != nil {
		panic(err)
	}
}

func Consume(channel *amqp.Channel, out chan amqp.Delivery) error {
	msgs, err := channel.Consume(
		QUEUE_PROCESSING,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}

	return nil
}
