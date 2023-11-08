package config

import (
	"encoding/json"
	"os"
)

func readJSON[T any](filename string) (T, error) {
	var config T

	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
