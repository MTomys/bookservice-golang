package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

}

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Display a list of the books on the reading list")
	}

	if r.Method == http.MethodPost {
		fmt.Fprintln(w, "Added a new book to the reading list")
	}
}

func (app *application) getUpdateDeleteBooksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			app.getBook(w, r)
		}
	case http.MethodPut:
		{
			app.updateBook(w, r)
		}
	case http.MethodDelete:
		{
			app.deleteBook(w, r)
		}
	default:
		{
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	}
}

func (app *application) getBook(w http.ResponseWriter, r *http.Request) {
	id := getIdFromRequest(w, r)
	fmt.Fprintf(w, "Display the details of book with id: %d", id)
}

func (app *application) updateBook(w http.ResponseWriter, r *http.Request) {
	id := getIdFromRequest(w, r)
	fmt.Fprintf(w, "Update the book with id: %d", id)
}

func (app *application) deleteBook(w http.ResponseWriter, r *http.Request) {
	id := getIdFromRequest(w, r)
	fmt.Fprintf(w, "Delete the book with id: %d", id)
}

func getIdFromRequest(w http.ResponseWriter, r *http.Request) int64 {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	return idInt
}
