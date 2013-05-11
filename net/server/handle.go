package server

import (
	"fmt"
	"github.com/rickbutton/speaker/net/client"
	"github.com/rickbutton/speaker/net/packet"
)

func (s *Server) handlePacket(p *packet.Packet, c *client.Client) {
	switch p.PacketType() {
	case packet.REQUEST:
		//send details packet
	case packet.ACK:
		//start streaming
	case packet.CLOSE:
		//close client
	case packet.PING:
		s.SendPacket(packet.NewPongPacket(0), c)
	case packet.PONG:
		//update client last seen time
	default:
		panic(fmt.Sprintf("Invalid packet type: %d", p.PacketType()))
	}
}

func (* Server) SendPacket(p *packet.Packet, c *client.Client) error {
  _, err = c.Conn.Write(p.Raw())
  return err
}
