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
	Enqueue(item interface{})
	Dequeue() interface{}
}

type TCPClient struct {
	Client
	conn     net.Conn
	packer   easytcp.Packer
	codec    easytcp.Codec
	messages []interface{}
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
		if len(c.messages) == 0 {
			return
		}

		log.Println("Writing...")
		msg := c.Dequeue()
		packedMsg, err := c.packer.Pack(msg.(*easytcp.Message))
		if err != nil {
			panic(err)
		}
		if _, err := c.conn.Write(packedMsg); err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}

func (c *TCPClient) Enqueue(item interface{}) {
	c.messages = append(c.messages, item)
}

func (c *TCPClient) Dequeue() interface{} {
	item := c.messages[0]
	c.messages = c.messages[1:]
	return item
}
