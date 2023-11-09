package server

import (
	"log"
	"rxjh-emu/public/config"
	"rxjh-emu/public/network"
	"sync"
)

type loginServer struct {
	config    config.ConfigLogin
	netPublic network.Listener
	netLocal  network.Listener

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
	// TODO add netLocal[easytcp] route list & middleware
	go ls.netLocal.Run(ls.config.LocalPort)

	ls.wg.Add(1)
	// TODO add netPublic[easytcp] route list & middleware
	go ls.netPublic.Run(ls.config.PublicPort)

	ls.wg.Wait()
}
