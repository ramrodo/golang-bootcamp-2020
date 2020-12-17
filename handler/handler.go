package handler

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ramrodo/golang-bootcamp-2020/model"
	"github.com/ramrodo/golang-bootcamp-2020/repository"
)

// Index - prints a hello message for API index
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// FilmIndex - returns all the films
func FilmIndex(w http.ResponseWriter, r *http.Request) {
	films, err := repository.FindAll()

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(films); err != nil {
		panic(err)
	}
}

// FilmShow - returns the film details requested
func FilmShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filmID := vars["filmID"]

	film, err := repository.Show(filmID)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(film); err != nil {
		panic(err)
	}
}

// FilmCreate - returns the film created
func FilmCreate(w http.ResponseWriter, r *http.Request) {
	var film model.Film

	// This is a good way to protect against malicious attacks on your server limiting JSON size
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &film); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // HTTP 422: unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t, err := repository.Create(film)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}

}
