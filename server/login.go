package server

import (
	"log"
	"rxjh-emu/public/config"
	"rxjh-emu/public/network"
	"sync"
)

type loginServer struct {
	config    config.ConfigLogin
	netPublic network.Server
	netLocal  network.Server

	wg *sync.WaitGroup
}

func NewLoginServer() *loginServer {
	return &loginServer{
		config:    config.LoadLoginConfig(),
		netPublic: *network.NewListener(),
		netLocal:  *network.NewListener(),

		wg: &sync.WaitGroup{},
	}
}

func (ls *loginServer) Start() {
	log.Println("Start LoginServer")

	ls.wg.Add(1)
	go ls.netLocal.Run(ls.config.LocalPort)

	ls.wg.Add(1)
	go ls.netPublic.Run(ls.config.PublicPort)

	ls.wg.Wait()
}
