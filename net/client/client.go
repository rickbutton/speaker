package client

import (
	"fmt"
	"github.com/rickbutton/speaker/net/decode"
	"net"
  "log"
  "os"
)

const (
  Port = 13332
)

var (
  logger *log.Logger = log.New(os.Stdout, "[net] ", log.LstdFlags)
)

type Client struct {
	Conn *net.TCPConn
  Id int
}

func NewClient(conn *net.TCPConn) *Client {
	c := new(Client)
	c.Conn = conn
	return c
}

func (c *Client) Listen(found chan bool) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", Port))
	if err != nil {
		panic(err)
	}
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
  found <- true
  close(found)
  listener.Close()
  c.Conn = conn.(*net.TCPConn)
	c.readFromServer()
}

func (c *Client) readFromServer() {
	buf := make([]byte, 4096)
	for {
		n, err := c.Conn.Read(buf)
		if err != nil {
      logger.Printf("Server disconnected")
      c.Shutdown()
      return
		}
		p := decode.DecodeRawPacket(buf[0:n])
		c.handlePacket(p)
	}
}

func Start() {
  logger.Printf("Starting client")
  client := NewClient(nil)
  found := make(chan bool)
	go client.FindServer(found)
	client.Listen(found)
}

func (c *Client) Shutdown() {
  c.Conn.Close()
}
