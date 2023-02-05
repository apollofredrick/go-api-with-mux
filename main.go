package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {

}
func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {

}
func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {

}
func DeletePersonEndpoint(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))

}
