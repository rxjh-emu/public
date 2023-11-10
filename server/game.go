package server

import (
	"log"
	"rxjh-emu/public/config"
	"rxjh-emu/public/network"
	"rxjh-emu/public/network/packer"
	"sync"

	"github.com/DarthPestilane/easytcp"
)

type gameServer struct {
	config config.ConfigGame

	localClient *network.TCPClient

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

	gs.wg.Add(1)
	cli := network.NewTCPClient(
		&packer.LocalPacker{},
		&easytcp.JsonCodec{},
	)
	gs.localClient = cli
	gs.localClient.Start()

	gs.wg.Wait()
}
