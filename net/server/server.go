package server

import (
  "net"
  "log"
  "os"
  "github.com/rickbutton/speaker/net/decode"
  "github.com/rickbutton/speaker/net/client"
)

var (
  logger *log.Logger = log.New(os.Stdout, "[net] ", log.LstdFlags)
  StreamPort = 13332
)

type Server struct {
  Clients []*client.Client
  lastId int
}

func NewServer() *Server {
  s := new(Server)
  s.Clients = make([]*client.Client, 0)
  return s
}

func (s *Server) OpenConn(ip net.IP) (*client.Client, error) {
  logger.Printf("Connecting to client %s:%d", ip.String(), StreamPort)
  socket, err := net.DialTCP("tcp4", nil, &net.TCPAddr{ ip, StreamPort })
  if err != nil {
    return nil, err
  }
  c := client.NewClient(socket)
  c.Id = s.lastId
  s.lastId++
  s.Clients = append(s.Clients, c)
  go s.readFromClient(c)
  logger.Printf("%d connected clients", len(s.Clients))
  return c, nil
}

func (s *Server) CloseConn(c *client.Client) {
  logger.Printf("Client %d at %s disconnected", c.Id, c.Conn.RemoteAddr().String())
  for i := range s.Clients {
    if s.Clients[i].Id == c.Id {
      //delete client from slice
      s.Clients[i] = s.Clients[len(s.Clients) - 1]
      s.Clients = s.Clients[0:len(s.Clients) - 1]
    }
  }
  c.Shutdown()
  logger.Printf("%d connected clients", len(s.Clients))
}

func (s *Server) readFromClient(c *client.Client) {
  buf := make([]byte, 4096)
  for {
    n, err := c.Conn.Read(buf)
    if (err != nil) {
      s.CloseConn(c)
      return
    }
    p := decode.DecodeRawPacket(buf[0:n])
    s.handlePacket(p, c)
  }
}

func Start() {
  s := NewServer()
  s.BroadcastListen()
}
