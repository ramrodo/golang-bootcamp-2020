package router

import (
	"net/http"

	"github.com/ramrodo/golang-bootcamp-2020/handler"
)

// Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes array
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handler.Index,
	},
	Route{
		"FilmIndex",
		"GET",
		"/films",
		handler.FilmIndex,
	},
	Route{
		"FilmShow",
		"GET",
		"/films/{filmID}",
		handler.FilmShow,
	},
	Route{
		"FilmCreate",
		"POST",
		"/films",
		handler.FilmCreate,
	},
}
