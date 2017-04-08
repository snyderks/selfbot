package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateResponse holds the json structure for Create()
type CreateResponse struct {
	Testval string `json:"test"`
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

// Edit edits a user's config
func Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["userID"]
	fmt.Fprint(w, "Config edited successfully\n"+user)
}

// Restart restarts a user's selfbot process
func Restart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["userID"]
	fmt.Fprint(w, "Process restarted successfully\n"+user)
}

// Start starts a user's selfbot process
func Start(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["userID"]
	fmt.Fprint(w, "Process started successfully\n"+user)
}

// Stop stops a user's selfbot process
func Stop(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["userID"]
	fmt.Fprint(w, "Process stopped successfully\n"+user)
}
