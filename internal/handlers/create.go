package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alysonandrade142/desafio-hitss/internal/model"
	"github.com/alysonandrade142/desafio-hitss/pkg/mq"
	"github.com/lithammer/shortuuid"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error on decode: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	uuid := shortuuid.New()
	body := model.QueueBody{
		MessageId: uuid,
		Method:    "CREATE",
		User:      user,
	}

	mq.Publish(r.Context(), body, mq.QUEUE_PROCESSING, uuid)

	response, err := mq.Consume(mq.QUEUE_RESPONSE, uuid)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Error on Consume: %v", err)
		return
	}

	responseMsg := fmt.Sprintf("User created with id: %v", response)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseMsg)
}
