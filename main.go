package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

type User struct {
	name string `json:"name"`
	mail string `json:"mail"`
	phone string `json:"phone"`
}

func main() {
	// Init Router
	router := max.NewRouter()

	// Create the Endpoints
	router.HandleFunc("/api/users", createUser, "POST")
	router.HandleFunc("/api/users/{id}", getUser, "GET")
	router.HandleFunc("/api/users", getUsers, "GET")

	// Launch the Server
	log.Fatal(http.ListenAndServe("0.0.0.0:50", router))
}
