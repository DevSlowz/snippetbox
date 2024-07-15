package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)

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
