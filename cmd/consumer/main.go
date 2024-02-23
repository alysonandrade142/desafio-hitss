package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/alysonandrade142/desafio-hitss/internal/model"
	"github.com/alysonandrade142/desafio-hitss/internal/repository"
	"github.com/alysonandrade142/desafio-hitss/pkg/configs"
	"github.com/alysonandrade142/desafio-hitss/pkg/mq"
	"github.com/streadway/amqp"
)

func main() {

	println("Starting server...")
	err := configs.Load()
	if err != nil {
		fmt.Printf("Error on load config: %v", err)
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		mq.QUEUE_PROCESSING, // Nome da fila
		false,               // Durable
		false,               // Delete when unused
		false,               // Exclusive
		false,               // No-wait
		nil,                 // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	msgs, err := ch.Consume(
		q.Name, // Nome da fila
		"",     // Consumer
		true,   // Auto-Ack
		false,  // Exclusive
		false,  // No-local
		false,  // No-Wait
		nil,    // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {

			var queueBody *model.QueueBody

			err := json.Unmarshal(d.Body, &queueBody)
			if err != nil {
				log.Printf("Cannot unmarshal queue body: %v", err)
				continue
			}

			var content interface{}

			switch queueBody.Method {
			case "CREATE":
				content, err = repository.Insert(queueBody.User)
				if err != nil {
					log.Printf("Cannot insert user: %v", err)
				}
			case "UPDATE":
				content, err = repository.Update(queueBody.ID, queueBody.User)
				if err != nil {
					log.Printf("Cannot update user: %v", err)
				}
			case "DELETE":
				content, err = repository.Delete(queueBody.ID)
				if err != nil {
					log.Printf("Cannot delete user: %v", err)
				}
			case "SEARCH":
				content, err = repository.Get(queueBody.ID)
				if err != nil {
					log.Printf("Cannot get user: %v", err)
				}
			default:
				content, err = repository.GetAll()
				if err != nil {
					log.Printf("Cannot get all users: %v", err)
				}
			}

			ctx := context.Background()
			responseBody := model.QueueBody{
				MessageId: queueBody.MessageId,
				Content:   content,
			}

			mq.Publish(ctx, responseBody, mq.QUEUE_RESPONSE, queueBody.MessageId)
		}
	}()

	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C\n")
	<-forever
}
