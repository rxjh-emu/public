package config

import (
	"errors"
	"log"
)

type ConfigGame struct {
	ServerId   int             `json:"server_id"`
	ServerName string          `json:"server_name"`
	LoginIp    string          `json:"login_ip"`
	LoginPort  int             `json:"login_port"`
	ServerIp   string          `json:"server_ip"`
	Database   ConfigDatabase  `json:"database"`
	Channels   []ConfigChannel `json:"channels"`
}

func LoadGameConfig() ConfigGame {
	config, err := readJSON[ConfigGame]("config_game.json")
	if err != nil {
		log.Println("Error reading JSON:", err)
		panic(errors.New("Error reading JSON: " + err.Error()))
	}
	return config
}
