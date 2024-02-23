package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alysonandrade142/desafio-hitss/internal/model"
	"github.com/alysonandrade142/desafio-hitss/pkg/mq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error on decode: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	conn, err := mq.NewRabbitMQ()
	if err != nil {
		log.Printf("Error on create connection: %v", err)
	}

	channel, err := mq.NewChannel(conn)
	if err != nil {
		log.Printf("Error on create channel: %v", err)
	}

	q, err := channel.QueueDeclare(
		mq.QUEUE_PROCESSING, // Nome da fila
		false,               // Durable
		false,               // Delete when unused
		false,               // Exclusive
		false,               // No-wait
		nil,                 // Arguments
	)

	if err != nil {
		log.Printf("Error on declare queue: %v", err)
	}

	err = channel.PublishWithContext(
		r.Context(),
		"",     // Exchange
		q.Name, // Key da fila
		false,  // Mandatory
		false,  // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("MEU PRIMEIRO TESTE - ALYSON"),
		})

	if err != nil {
		log.Printf("Error on publish: %v", err)
	}

	// id, err := repository.Insert(user)

	var resp map[string]interface{}

	// if err != nil {
	// 	resp = map[string]interface{}{
	// 		"Error": "A exception occurred: " + err.Error(),
	// 	}
	// } else {
	// 	resp = map[string]interface{}{
	// 		"Message": fmt.Sprintf("User created with id: %d", id),
	// 	}
	// }

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func Publish(w http.ResponseWriter, r *http.Request) {
}

func Receiver(w http.ResponseWriter, r *http.Request) {

}
