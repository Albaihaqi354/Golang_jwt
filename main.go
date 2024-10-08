package main

import (
	"golang_jwt/configs"
	"golang_jwt/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	configs.ConnectDb()

	r := mux.NewRouter()

	router := r.PathPrefix("/api").Subrouter()

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	log.Println("Server Running on Port: 8080")
	http.ListenAndServe(":8080", router)
}
