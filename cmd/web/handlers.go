package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Home Handler function which writes a byte slice containg Hello from Snippetbox in the response body

func home(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Server", "Go")

	w.Write([]byte("Hello from Snippetbox"))

}

// Controller/handler for Viewing a snippet

func snippetView(w http.ResponseWriter, r *http.Request) {

	// Extract value from id wildcard if its invalid or less than 1 return not found response

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {

		http.NotFound(w, r)

		return

	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)

	// w.Write([]byte(msg))

}

// Controller for creating a snippet

func snippetCreate(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Display a form for creating a new snippet"))

}

// Handler for saving a snippet

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	// Send 201 status code

	w.WriteHeader(http.StatusCreated)

	// Write response body

	w.Write([]byte("Save a new snippet..."))

}
