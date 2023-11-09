package main

import (
	"flag"
	"rxjh-emu/public/server"

	"log"
)

var typePtr *string

func init() {
	typePtr = flag.String("type", "", "Denotes what type of server to start: login, world, channel")
	flag.Parse()
}

func main() {
	switch *typePtr {
	case "login":
		s := server.NewLoginServer()
		s.Start()
	case "game":
		s := server.NewGameServer()
		s.Start()
	default:
		if *typePtr != "" {
			log.Fatalf("Unkown server type: %s", *typePtr)
		} else {
			log.Fatal("Denotes what type of server to start: login, game")
		}
	}
}
