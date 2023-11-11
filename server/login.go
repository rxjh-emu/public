package server

import (
	"log"
	"rxjh-emu/public/config"
	"rxjh-emu/public/network"
	"rxjh-emu/public/network/packer"
	localRoutes "rxjh-emu/public/network/routes/local"
	"rxjh-emu/public/utils/consts/opcodes"
	"sync"

	"github.com/DarthPestilane/easytcp"
)

type loginServer struct {
	config   config.ConfigLogin
	netLocal network.Listener

	localRoutes map[uint16]network.RouteInfo

	wg *sync.WaitGroup
}

func NewLoginServer() *loginServer {
	return &loginServer{
		config: config.LoadLoginConfig(),
		netLocal: *network.NewListener(
			&packer.LocalPacker{},
			&easytcp.JsonCodec{},
		),

		localRoutes: make(map[uint16]network.RouteInfo),

		wg: &sync.WaitGroup{},
	}
}

func (ls *loginServer) Start() {
	log.Println("Start LoginServer")

	ls.initRoutes()

	ls.wg.Add(1)
	// TODO add netLocal[easytcp] route list & middleware
	ls.netLocal.RegisterRoutes(ls.localRoutes)
	go ls.netLocal.Run(ls.config.LocalPort)

	ls.wg.Wait()
}

func (ls *loginServer) initRoutes() {
	// LOCAL ROUTES
	ls.localRoutes[opcodes.L_REQ_PING] = &localRoutes.L_REQ_PING{}
	ls.localRoutes[opcodes.L_REQ_SERVER_REGISTER] = &localRoutes.L_REQ_SERVER_REGISTER{}

	// CLIENT ROUTES
}
