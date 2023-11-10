package server

import (
	"log"
	"rxjh-emu/public/config"
	"sync"
)

type gameServer struct {
	config config.ConfigGame

	wg *sync.WaitGroup
}

func NewGameServer() *gameServer {
	c := config.LoadGameConfig()

	return &gameServer{
		config: c,

		wg: &sync.WaitGroup{},
	}
}

func (gs *gameServer) initChannels() {
	// for _, chn := range gs.config.Channels {
	// 	gs.wg.Add(1)
	// 	s := network.NewListener()
	// 	// TODO add channel[easytcp] route list & middleware
	// 	go s.Run(chn.Port)
	// }
}

func (gs *gameServer) Start() {
	log.Println("Start GameServer")

	gs.initChannels()

	gs.wg.Wait()
}
