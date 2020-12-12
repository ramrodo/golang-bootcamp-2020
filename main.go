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
	// csvFile, err := os.Open(config.C.Database.File)
	// defer csvFile.Close()

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// reader := csv.NewReader(bufio.NewReader(csvFile))

	// all, err := repository.FindAll(reader)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	router := router.NewRouter()

	fmt.Printf("Listening on:%s:%s\n", config.C.Server.URL, config.C.Server.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.C.Server.Port), router))
}
