package config

import (
	"encoding/json"
	"io/ioutil"

	"log"
)

type Config struct {
	NoOfFloors            int `json:"no_of_floors"`
	MainCorridorsPerFloor int `json:"main_corridors_per_floor"`
	SubCorridorsPerFloor  int `json:"sub_corridors_per_floor"`
}

func Get() *Config {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("Error reading input file", err)
	}

	i := new(Config)
	err = json.Unmarshal(file, i)
	if err != nil {
		log.Fatal("Error parsing yml file", err)
	}

	return i
}
