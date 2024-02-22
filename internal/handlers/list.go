package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/alysonandrade142/desafio-hitss/internal/repository"
	"github.com/gorilla/mux"
)

func List(w http.ResponseWriter, r *http.Request) {
	users, err := repository.GetAll()
	if err != nil {
		log.Printf("Error on get all users: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Printf("Error on get id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	user, err := repository.Get(int64(id))
	if err != nil {
		log.Printf("Error on get user: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
