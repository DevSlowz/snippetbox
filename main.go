package main

import (
	"log"
	"net/http"
)

// Home Handler function which writes a byte slice containg Hello from Snippetbox in the response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

// Controller for Viewing a snippet
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// Controller for creating a snippet
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet"))
}

func main() {

	// http.NewServeMux initializes a new sermux(web server)
	mux := http.NewServeMux()
	// register home function as the handler for "/" URL pattern
	mux.HandleFunc("/", home)

	// View for viewing snippets
	mux.HandleFunc("/snippet/view", snippetView)
	// View for creating snippets
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("starting server on : 4000")

	// http.ListenAndServe function starts a new web server
	// Pass in the TCP network address to listen on port 4000 and the servemux created(mux)
	// If http.ListenAndServe returns an error log.fatal() will log the error and exit
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
