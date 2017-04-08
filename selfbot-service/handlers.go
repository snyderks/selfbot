package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateResponse holds the json structure for Create()
type CreateResponse struct {
}

// Index responds with Hello World to view if server is running
func Index(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["userID"]
	fmt.Fprint(w, "Hello world\n"+user)
}

// Create creates a new selfbot if one is not created already
func Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["userID"]
	fmt.Fprint(w, "Bot created successfully\n"+user)
}

// Edit edit's a user's config
func Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["userID"]
	fmt.Fprint(w, "Config edited successfully\n"+user)
}

// Restart restart's a user's selfbot process
func Restart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["userID"]
	fmt.Fprint(w, "Process restarted successfully\n"+user)
}
