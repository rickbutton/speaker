package client

import (
	"github.com/rickbutton/speaker/net/packet"
)

const (
	port = 13331
)

var (
	broadcast net.IP = net.IPv4(255, 255, 255, 255)
)

func FindServer() {
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{broadcast, port})
	defer socket.Close()
	if err != nil {
		panic(err)
	}
	p := packet.NewPingPacket(0)
	socket.Write(p.Raw())
}
