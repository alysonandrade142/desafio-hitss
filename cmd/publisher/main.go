package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"db-queue", // Nome da fila
		false,      // Durable
		false,      // Delete when unused
		false,      // Exclusive
		false,      // No-wait
		nil,        // Arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello, TESTE ALYSON!"

	err = ch.Publish(
		"",     // Exchange
		q.Name, // Key da fila
		false,  // Mandatory
		false,  // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C\n")

	fmt.Printf(" [x] Sent %s\n", body)
}
