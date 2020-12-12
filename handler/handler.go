package handler

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"

	"gopkg.in/resty.v1"

	"github.com/gorilla/mux"
	"github.com/ramrodo/golang-bootcamp-2020/model"
	"github.com/ramrodo/golang-bootcamp-2020/repository"
)

// Index .
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// FilmIndex .
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

// FilmShow .
func FilmShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filmID := vars["filmID"]

	client := resty.New()

	var film model.Film
	resp, err := client.R().Get(fmt.Sprintf("https://ghibliapi.herokuapp.com/films/%s", filmID))
	json.Unmarshal(resp.Body(), &film)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(resp.StatusCode())
	if err := json.NewEncoder(w).Encode(film); err != nil {
		panic(err)
	}
}

// FilmCreate .
func FilmCreate(w http.ResponseWriter, r *http.Request) {
	var film model.Film

	// Notice that we use io.LimitReader.
	// This is a good way to protect against malicious attacks on your server.
	// Imagine if someone wanted to send you 500GBs of json!
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

	// CreateFilm from repository
	t := repository.CreateFilm(film)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}

}
