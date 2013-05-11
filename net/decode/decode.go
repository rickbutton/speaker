package decode

import (
	"encoding/binary"
	"fmt"
	"github.com/rickbutton/speaker/net/packet"
)

func DecodeRawPacket(raw []byte) packet.Packet {
	t := packet.PacketType(raw[1])
	switch t {
	case packet.REQUEST, packet.ACK, packet.CLOSE, packet.PING, packet.PONG:
		return packet.NewSimplePacket(raw[0], t)
	case packet.DETAILS:
		return packet.NewDetailsPacket(raw[0],
			packet.SampleRate(raw[2]),
			packet.SampleResolution(raw[3]))
	case packet.STREAM:
		timestamp := binary.LittleEndian.Uint64(raw[2:10])
		return packet.NewStreamPacket(raw[0], timestamp, raw[10:])
	}
	panic(fmt.Sprintf("Invalid packet with data %v", raw))
}
