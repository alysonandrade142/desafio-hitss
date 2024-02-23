package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alysonandrade142/desafio-hitss/internal/model"
	"github.com/alysonandrade142/desafio-hitss/pkg/mq"
	uuid "github.com/satori/go.uuid"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error on decode: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	uuid := uuid.NewV4()
	body := model.QueueBody{
		MessageId: uuid,
		Method:    "CREATE",
		User:      user,
	}

	mq.Publish(r.Context(), body, mq.QUEUE_PROCESSING)

	response := mq.Consume(mq.QUEUE_RESPONSE, uuid)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
