package client

import (
	"fmt"
	"github.com/rickbutton/speaker/net/decode"
	"github.com/rickbutton/speaker/net/server"
	"net"
)

type Client struct {
	Conn *net.TCPConn
}

func NewClient(conn *net.TCPConn) *Client {
	c := new(Client)
	c.Conn = conn
	return c
}

func Listen() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", server.Port))
	if err != nil {
		panic(err)
	}
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	client := NewClient(conn)
	go readFromServer(client)
}

func readFromServer(c *Client) {
	buf := make([]byte, 4096)
	for {
		n, err := c.Conn.Read(buf)
		if err != nil {
			panic(err)
		}
		p := decode.DecodeRawPacket(buf[0:n])
		c.handlePacket(&p)
	}
}

func Start() {
	go FindServer()
	go Listen()
}
