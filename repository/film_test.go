package repository

import (
	"testing"

	"github.com/ramrodo/golang-bootcamp-2020/config"
	"github.com/spf13/viper"
)

func TestShow(t *testing.T) {
	viper.AddConfigPath("../")
	config.ReadConfig()

	film, _ := Show("2baf70d1-42bb-4437-b551-e5fed5a87abe")

	if film.Title != "Castle in the Sky" {
		t.Error("Error at retrieving film data from upstream API")
	}
}
