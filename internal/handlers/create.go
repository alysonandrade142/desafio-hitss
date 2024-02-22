package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alysonandrade142/desafio-hitss/internal/model"
	"github.com/alysonandrade142/desafio-hitss/internal/repository"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error on decode: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := repository.Insert(user)

	var resp map[string]interface{}

	if err != nil {
		resp = map[string]interface{}{
			"Error": "A exception occurred: " + err.Error(),
		}
	} else {
		resp = map[string]interface{}{
			"Message": fmt.Sprintf("User created with id: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
