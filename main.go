package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ramrodo/golang-bootcamp-2020/config"
	"github.com/ramrodo/golang-bootcamp-2020/router"
)

func main() {
	config.ReadConfig()

	router := router.NewRouter()

	fmt.Printf("Listening on: %s:%s\n", config.C.Server.URL, config.C.Server.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.C.Server.Port), router))
}
