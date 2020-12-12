package repository

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ramrodo/golang-bootcamp-2020/config"
	"github.com/ramrodo/golang-bootcamp-2020/model"
	"gopkg.in/resty.v1"
)

// FindAll - reads CSV file and returns an array of Films
func FindAll() ([]model.Film, error) {
	csvFile, err := os.Open(config.C.Database.File)
	defer csvFile.Close()

	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))

	data, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	films := []model.Film{}

	for i, row := range data {
		if i == 0 {
			continue
		}

		films = append(films, model.Film{
			ID:          row[0],
			Title:       row[1],
			Description: row[2],
			Director:    row[3],
			Producer:    row[4],
			ReleaseDate: row[5],
			RtScore:     row[6],
		})
	}

	return films, nil
}

// Create .
func Create(film model.Film) (model.Film, error) {
	// FIXME: How to share the reader between handlers/repository functions
	csvFile, err := os.OpenFile(config.C.Database.File, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer csvFile.Close()

	if err != nil {
		log.Fatalln(err)
	}

	w := csv.NewWriter(csvFile)
	defer w.Flush()

	if err := w.Write([]string{film.ID, film.Title, film.Description, film.Director, film.Producer, film.ReleaseDate, film.RtScore}); err != nil {
		log.Fatalln("error writing record to file", err)
		return model.Film{}, err
	}

	return film, nil
}

// Show .
func Show(filmID string) (model.Film, error) {
	client := resty.New()

	var film model.Film
	resp, err := client.R().Get(fmt.Sprintf("https://ghibliapi.herokuapp.com/films/%s", filmID))
	json.Unmarshal(resp.Body(), &film)

	if err != nil {
		return film, err
	}

	return film, nil
}
