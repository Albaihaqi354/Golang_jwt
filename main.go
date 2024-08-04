package main

import (
	"golang_jwt/configs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	configs.ConnectDb()

	r := mux.NewRouter()

	router := r.PathPrefix("/api").Subrouter()

	log.Println("Server Running on Port: 8080")
	http.ListenAndServe(":8080", router)
}
