package server

import (
	"github.com/rickbutton/speaker/net/decode"
	"github.com/rickbutton/speaker/net/packet"
	"net"
)

const (
	port = 13331
)

var (
	broadcast net.IP = net.IPv4(255, 255, 255, 255)
)

func (s *Server) BroadcastListen() {
	socket, err := net.ListenUDP("udp4", nil, &net.UDPAddr{broadcast, port})
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
			client, err := server.OpenConn(addr.IP)
			if err != nil {
				panic(err)
			}
			s.handlePacket(p, client)
		}
	}
}
