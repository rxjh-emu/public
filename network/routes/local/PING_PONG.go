package routes

import (
	"log"
	"rxjh-emu/public/network"
	"rxjh-emu/public/utils/consts/opcodes"

	"github.com/DarthPestilane/easytcp"
)

type L_REQ_PING struct {
	network.RouteInfo

	Time int64 `json:"time"`
}

type L_RES_PONG struct {
	Time int64 `json:"time"`
}

func (p *L_REQ_PING) HandleRequest(c easytcp.Context) {
	log.Println("REQ PING Read")
	var data L_REQ_PING
	_ = c.Bind(&data)

	log.Printf("time: %d", data.Time)

	p.Response(c)
}

func (p *L_REQ_PING) Response(c easytcp.Context) {
	data := &L_RES_PONG{}
	data.Time = p.Time
	c.SetResponse(opcodes.L_RES_PONG, data)
}
