package routes

import (
	"log"
	"rxjh-emu/public/config"
	"rxjh-emu/public/network"

	"github.com/DarthPestilane/easytcp"
)

type L_REQ_SERVER_REGISTER struct {
	network.RouteInfo `json:"-"`

	ServerId   int                    `json:"server_id"`
	ServerName string                 `json:"server_name"`
	ServerIp   string                 `json:"server_ip"`
	Channels   []config.ConfigChannel `json:"channels"`
}

type L_RES_SERVER_REGISTER struct {
	Result bool `json:"result"`
}

func (p *L_REQ_SERVER_REGISTER) HandleRequest(c easytcp.Context) {
	// log.Printf("0x000A(10), Data: %s", c.Request().Data())
	// TODO register service
	var data L_REQ_SERVER_REGISTER
	_ = c.Bind(&data)

	log.Printf("ServerName: %s", data.ServerName)
}

func (p *L_REQ_SERVER_REGISTER) Responses() {

}
