package main

import (
	"fmt"
	"net/http"

	"github.com/alysonandrade142/desafio-hitss/internal/handlers"
	"github.com/alysonandrade142/desafio-hitss/pkg/configs"
	"github.com/gorilla/mux"
)

func main() {
	println("Starting server...")
	err := configs.Load()
	if err != nil {
		fmt.Printf("Error on load config: %v", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.List).Methods("GET")
	r.HandleFunc("/users", handlers.List).Methods("GET")
	r.HandleFunc("/getUser/{id:[0-9]+}", handlers.GetUser).Methods("GET")

	r.HandleFunc("/createUser", handlers.Create).Methods("POST")
	r.HandleFunc("/updateUser/{id:[0-9]+}", handlers.Update).Methods("PUT")
	r.HandleFunc("/deleteUser/{id:[0-9]+}", handlers.Delete).Methods("DELETE")

	fmt.Println("Server running on port", configs.GetServerPort())
	http.ListenAndServe(":"+configs.GetServerPort(), r)
}
