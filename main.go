package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Home Handler function which writes a byte slice containg Hello from Snippetbox in the response body
func home(w http.ResponseWriter, r *http.Request) {
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
	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

// Controller for creating a snippet
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet"))
}

// Handler for saving a snippet
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new snippet..."))
}

func main() {

	// http.NewServeMux initializes a new sermux(web server)
	mux := http.NewServeMux()
	// register home function as the handler for "/" URL pattern
	mux.HandleFunc("GET /{$}", home) // This route will only respond to exacy matches on / only

	// View for viewing snippets
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	// View for creating snippets
	mux.HandleFunc("GET /snippet/create", snippetCreate)

	// View for create and save a snipped to the database
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on : 4000")

	// http.ListenAndServe function starts a new web server
	// Pass in the TCP network address to listen on port 4000 and the servemux created(mux)
	// If http.ListenAndServe returns an error log.fatal() will log the error and exit
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
