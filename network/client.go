package network

import (
	"fmt"
	"log"
	"net"
	"rxjh-emu/public/utils/consts/opcodes"
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

func NewTCPClient(loginIp string, loginPort int, packer easytcp.Packer, codec easytcp.Codec) *TCPClient {

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", loginIp, loginPort))
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
	for {
		go c.Reading()
		go c.Writing()
		time.Sleep(time.Millisecond * 100)
	}
}

func (c *TCPClient) Reading() {
	msg, err := c.packer.Unpack(c.conn)
	if err != nil {
		panic(err)
	}
	fullSize := msg.MustGet("fullSize")
	log.Printf("ack received | fullSize:(%d) id:(%v) dataSize:(%d) data: %s", fullSize, msg.ID(), len(msg.Data()), msg.Data())
}

func (c *TCPClient) Writing() {
	if len(c.messages) == 0 {
		return
	}

	msg := c.Dequeue()
	packedMsg, err := c.packer.Pack(msg.(*easytcp.Message))
	if err != nil {
		panic(err)
	}
	if _, err := c.conn.Write(packedMsg); err != nil {
		panic(err)
	}
}

func (c *TCPClient) Enqueue(opcode uint16, packet interface{}) {
	// log.Printf("Enqueue opcode: %d, packet: %s", opcode, packet.(RouteInfo))
	data, err := c.codec.Encode(packet)
	if err != nil {
		panic(err)
	}
	msg := easytcp.NewMessage(opcode, data)
	c.messages = append(c.messages, msg)
}

func (c *TCPClient) Dequeue() interface{} {
	item := c.messages[0]
	c.messages = c.messages[1:]
	return item
}

func (c *TCPClient) RegisterServerToLoginServer(data interface{}) {
	log.Println("Established to LoginServer...")
	c.Enqueue(opcodes.L_REQ_SERVER_REGISTER, data)
}
