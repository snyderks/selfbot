package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateResponse holds the json structure for Create()
type CreateResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type CreateRequest struct {
	Token string `json:"Token"`
}

// Index responds with Hello World to view if server is running
func Index(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["userID"]
	fmt.Fprint(w, "Hello world\n"+user)
}

// Create creates a new selfbot if one is not created already
func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var t CreateRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, t)
}

// Delete deletes a selfbot
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["userID"]
	fmt.Fprint(w, "Bot deleted successfully\n"+user)
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
