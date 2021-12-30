package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")

	err := http.ListenAndServe(":8081", router)
	if err != nil {
		panic(err)
	}
}

func main() {
	initializeMigration()
	initializeRouter()
}
