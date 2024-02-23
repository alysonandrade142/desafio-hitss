package mq

import (
	"context"
	"encoding/json"
	"log"

	"github.com/alysonandrade142/desafio-hitss/internal/model"
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

func Consume(queue string, uuid string) (interface{}, error) {

	var out chan amqp.Delivery
	conn, err := NewRabbitMQ()
	if err != nil {
		log.Printf("Error on create connection: %v", err)
		return nil, err
	}

	defer CloseConnection(conn)

	channel, err := NewChannel(conn)
	if err != nil {
		log.Printf("Error on create channel: %v", err)
		return nil, err
	}
	defer CloseChannel(channel)

	msgs, err := channel.Consume(
		queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Printf("Error on consume: %v", err)
		return nil, err
	}

	for msg := range msgs {

		var body model.QueueBody
		err := json.Unmarshal(msg.Body, &body)
		if err != nil {
			log.Printf("Cannot unmarshaly: %v", err)
			return nil, err
		}

		if body.MessageId == uuid {
			channel.Ack(msg.DeliveryTag, false)
			return body.Content, nil
		}
		out <- msg
	}

	return nil, nil
}

func Publish(ctx context.Context, pBody interface{}, queue string, uuid string) {

	body, err := json.Marshal(pBody)
	if err != nil {
		log.Printf("Error on marshal pBody: %v", err)
	}

	conn, err := NewRabbitMQ()
	if err != nil {
		log.Printf("Error on create connection: %v", err)
	}
	defer CloseConnection(conn)

	channel, err := NewChannel(conn)
	if err != nil {
		log.Printf("Error on create channel: %v", err)
	}
	defer CloseChannel(channel)

	q, err := channel.QueueDeclare(
		queue,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Printf("Error on declare queue: %v", err)
	}

	err = channel.PublishWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		log.Printf("Error on publish: %v", err)
	}
}
