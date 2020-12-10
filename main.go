// package main

// import (
// 	"bufio"
// 	"encoding/csv"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/ramrodo/golang-bootcamp-2020/config"
// 	"github.com/ramrodo/golang-bootcamp-2020/usecase/repository"
// )

// func main() {
// 	config.ReadConfig()
// 	csvFile, err := os.Open(config.C.Database.File)
// => Add defer

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	reader := csv.NewReader(bufio.NewReader(csvFile))

// 	all, err := repository.FindAll(reader)

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	fmt.Println("Title:", all[0].Title)
// 	fmt.Println("Description:", all[0].Description)
// }
