package main

import (
	"log"
	"net/http"
)

// Home Handler function which writes a byte slice containg Hello from Snippetbox in the response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("starting server on : 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
