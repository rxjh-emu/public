package network

import (
	"log"
	"net"
	"time"

	"github.com/DarthPestilane/easytcp"
)

type Client interface {
	Start()
	Reading()
	Writing()
}

type TCPClient struct {
	conn   net.Conn
	packer easytcp.Packer
	codec  easytcp.Codec
}

func NewTCPClient(packer easytcp.Packer, codec easytcp.Codec) *TCPClient {

	conn, err := net.Dial("tcp", "127.0.0.1:12100")
	if err != nil {
		panic(err)
	}

	return &TCPClient{
		conn:   conn,
		packer: packer,
		codec:  codec,
	}
}

func (c *TCPClient) Start() {
	go c.Reading()
	go c.Writing()
}

func (c *TCPClient) Reading() {
	for {
		// TODO READ
		log.Println("Reading...")
		time.Sleep(time.Second)
	}
}

func (c *TCPClient) Writing() {
	for {
		// TODO SEND
		log.Println("Writing...")
		time.Sleep(time.Second)
	}
}
