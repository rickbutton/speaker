package server

import (
	"github.com/rickbutton/speaker/net/decode"
	"github.com/rickbutton/speaker/net/packet"
	"net"
  _ "log"
)

const (
	port = 13331
)

var (
	listen net.IP = net.IPv4(0, 0, 0, 0)
)

func (s *Server) BroadcastListen() {
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{listen, port})
  logger.Printf("Starting auto discovery server")
	if err != nil {
		panic(err)
	}
	for {
		data := make([]byte, 1024)
		n, addr, err := socket.ReadFromUDP(data)
		if err != nil {
			panic(err)
		}
		p := decode.DecodeRawPacket(data[0:n])
		if p.PacketType() == packet.PING {
			client, err := s.OpenConn(addr.IP)
			if err != nil {
				panic(err)
			}
			s.handlePacket(p, client)
		}
	}
}
