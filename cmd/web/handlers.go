package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Home Handler function which writes a byte slice containg Hello from Snippetbox in the response body

func home(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	// Read template file into template set use the spread operator to pass all files
	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write template content as the response body
	// ExecuteTemplate writes the content of the "base" template as the response body
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	// w.Write([]byte("Hello from Snippetbox"))

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
