package server

import (
  "net"
  "github.com/rickbutton/speaker/net/decode"
  "github.com/rickbutton/speaker/net/client"
)

const (
  Port = 13332
)

type Server struct {
  Clients []*client.Client
}

func NewServer() *Server {
  s := new(Server)
  s.Clients = make([]*client.Client, 0)
  return s
}

func (s *Server) OpenConn(ip net.IP) (*client.Client, error) {
  socket, err := net.DialTCP("tcp4", nil, &net.TCPAddr{ ip, Port })
  if err != nil {
    return nil, err
  }
  c := client.NewClient(socket)
  s.Clients = append(s.Clients, c)
  go s.readFromClient(c)
  return c, nil
}

func (s *Server) readFromClient(c *client.Client) {
  buf := make([]byte, 4096)
  for {
    n, err := c.Conn.Read(buf)
    if (err != nil) {
      panic(err)
    }
    p := decode.DecodeRawPacket(buf[0:n])
    s.handlePacket(&p, c)
  }
}

func Start() *Server {
  s := NewServer()
  go s.BroadcastListen()
  return s
}
