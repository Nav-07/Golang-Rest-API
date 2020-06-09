package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type User struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
	Phone string `json:"phone"`
}

var users []User

func createUser(h http.ResponseWriter, r *http.Request) {
}
func getUser(h http.ResponseWriter, r *http.Request) {
}
func getUsers(h http.ResponseWriter, r *http.Request) {
	h.Header().Set("Content-Type", "application/json")
	json.NewEncoder(h).Encode(users)
}

func main() {
	// Init Router
	router := mux.NewRouter()

	// Enter Mock Data
	users = append(users, User{Name: "Nav Pahwa", Mail: "donkey@gmail.com", Phone: "+1000000000"})
	users = append(users, User{Name: "John Doe", Mail: "john.doe@gmail.com", Phone: "+1000654000"})
	users = append(users, User{Name: "Sam Smith", Mail: "sam.smith@gmail.com", Phone: "+1007830000"})

	// Create the Endpoints
	router.HandleFunc("/api/users", createUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/api/users", getUsers).Methods("GET")

	// Launch the Server
	log.Fatal(http.ListenAndServe("0.0.0.0:50", router))
}
