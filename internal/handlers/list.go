package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/alysonandrade142/desafio-hitss/internal/model"
	"github.com/alysonandrade142/desafio-hitss/pkg/mq"
	"github.com/gorilla/mux"
	"github.com/lithammer/shortuuid"
)

func List(w http.ResponseWriter, r *http.Request) {

	uuid := shortuuid.New()
	body := model.QueueBody{
		MessageId: uuid,
		Method:    "LIST",
	}
	mq.Publish(r.Context(), body, mq.QUEUE_PROCESSING, uuid)

	response, err := mq.Consume(mq.QUEUE_RESPONSE, uuid)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Error on Consume: %v", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Printf("Error on get id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	uuid := shortuuid.New()

	body := model.QueueBody{
		MessageId: uuid,
		ID:        int64(id),
		Method:    "SEARCH",
	}

	mq.Publish(r.Context(), body, mq.QUEUE_PROCESSING, uuid)

	response, err := mq.Consume(mq.QUEUE_RESPONSE, uuid)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Error on Consume: %v", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
