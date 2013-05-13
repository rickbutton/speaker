package client

import (
	"github.com/rickbutton/speaker/net/packet"
)

func (c *Client) handlePacket(p packet.Packet) {
	switch p.PacketType() {
	case packet.DETAILS:
		//handle stream details
	case packet.STREAM:
		//handle streaming
	case packet.CLOSE:
		//close client
	case packet.PING:
		c.SendPacket(packet.NewPongPacket(0))
	case packet.PONG:
		//handle server last seen time
	default:
		logger.Printf("Invalid packet type: %d", p.PacketType())
	}
}

func (c *Client) SendPacket(p packet.Packet) error {
  _, err := c.Conn.Write(p.Raw())
	return err
}
