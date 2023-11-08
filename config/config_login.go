package config

import (
	"errors"
	"log"
)

type ConfigLogin struct {
	PublicIp   string         `json:"public_ip"`
	PublicPort int            `json:"public_port"`
	LocalIp    string         `json:"local_ip"`
	LocalPort  int            `json:"local_port"`
	Database   ConfigDatabase `json:"database"`
}

func LoadLoginConfig() ConfigLogin {
	config, err := readJSON[ConfigLogin]("config_login.json")
	if err != nil {
		log.Println("Error reading JSON:", err)
		panic(errors.New("Error reading JSON: " + err.Error()))
	}
	return config
}
