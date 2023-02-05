package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

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
	people = append(people, Person{ID: "1", Firstname: "Nick", Lastname: "Pope", Address: &Address{City: "Dublin", State: "California"}})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))

}
