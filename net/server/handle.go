package server

import (
	"github.com/rickbutton/speaker/net/client"
	"github.com/rickbutton/speaker/net/packet"
)

func (s *Server) handlePacket(p packet.Packet, c *client.Client) {
	switch p.PacketType() {
	case packet.REQUEST:
		//send details packet
	case packet.ACK:
		//start streaming
	case packet.CLOSE:
    s.CloseConn(c)
	case packet.PING:
		s.SendPacket(packet.NewPongPacket(0), c)
	case packet.PONG:
    //do nothing, client keep alive
	default:
		logger.Printf("Invalid packet with type: %d", p.PacketType())
	}
}

func (* Server) SendPacket(p packet.Packet, c *client.Client) error {
  _, err := c.Conn.Write(p.Raw())
  return err
}
