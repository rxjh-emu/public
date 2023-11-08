package server

import (
	"log"
	"rxjh-emu/public/network"
	"sync"
)

type loginServer struct {
	netServer network.Server

	wg *sync.WaitGroup
}

func NewLoginServer() *loginServer {
	return &loginServer{
		netServer: *network.NewListener(),

		wg: &sync.WaitGroup{},
	}
}

func (ls *loginServer) Start() {
	log.Println("Start LoginServer")

	ls.wg.Add(1)
	go ls.netServer.Run()

	ls.wg.Wait()
}
