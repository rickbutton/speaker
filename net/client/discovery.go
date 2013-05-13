package client

import (
	"github.com/rickbutton/speaker/net/packet"
  "net"
  "time"
)

const (
	port = 13331
)

var (
	broadcast net.IP = net.IPv4(255, 255, 255, 255)
)

func (c *Client) FindServer(found chan bool) {
  logger.Printf("Trying to find server")
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{broadcast, port})
  defer socket.Close()
	if err != nil {
		panic(err)
	}
	p := packet.NewPingPacket(0)
  socket.Write(p.Raw())
  for {
    select {
    case <-found:
      logger.Printf("Found server, turning off auto discovery")
      return
    case <-time.After(2 * time.Second):
      logger.Printf("Still looking for server...")
      socket.Write(p.Raw())
    }
  }
}
