package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/alysonandrade142/desafio-hitss/internal/model"
	"github.com/alysonandrade142/desafio-hitss/pkg/mq"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

func Delete(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Printf("Error on get id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	uuid := uuid.NewV4()
	body := model.QueueBody{
		MessageId: uuid,
		Method:    "DELETE",
		ID:        int64(id),
	}

	mq.Publish(r.Context(), body, mq.QUEUE_PROCESSING)

	response := mq.Consume(mq.QUEUE_RESPONSE, uuid)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
