package server

import (
	"log"
	"rxjh-emu/public/config"
	"rxjh-emu/public/network"
	"rxjh-emu/public/network/packer"
	routes "rxjh-emu/public/network/routes/local"
	"sync"

	"github.com/DarthPestilane/easytcp"
)

type gameServer struct {
	config config.ConfigGame

	localClient *network.TCPClient

	wg *sync.WaitGroup
}

func NewGameServer() *gameServer {
	cfg := config.LoadGameConfig()

	return &gameServer{
		config: cfg,
		localClient: network.NewTCPClient(
			cfg.LoginIp,
			cfg.LoginPort,
			&packer.LocalPacker{},
			&easytcp.JsonCodec{},
		),

		wg: &sync.WaitGroup{},
	}
}

// func (gs *gameServer) initChannels() {
// 	// for _, chn := range gs.config.Channels {
// 	// 	gs.wg.Add(1)
// 	// 	s := network.NewListener()
// 	// 	// TODO add channel[easytcp] route list & middleware
// 	// 	go s.Run(chn.Port)
// 	// }
// }

func (gs *gameServer) Start() {
	log.Println("Start GameServer")

	// gs.initChannels()

	gs.wg.Add(1)
	go gs.localClient.Start()

	gs.RegisterServerToLoginServer()

	gs.wg.Wait()
}

func (gs *gameServer) RegisterServerToLoginServer() {
	data := routes.L_REQ_SERVER_REGISTER{
		ServerId:   gs.config.ServerId,
		ServerName: gs.config.ServerName,
		ServerIp:   gs.config.ServerIp,
		Channels:   gs.config.Channels,
	}
	gs.localClient.RegisterServerToLoginServer(data)
}
