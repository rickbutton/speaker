package packet

import (
	"encoding/binary"
)

type PacketType uint8
type SampleRate uint8
type SampleResolution uint8

const (
	REQUEST PacketType = 0
	DETAILS PacketType = 1
	ACK     PacketType = 2
	STREAM  PacketType = 3
	CLOSE   PacketType = 4
	PING    PacketType = 5
	PONG    PacketType = 6

	R8000  SampleRate = 0
	R11025 SampleRate = 1
	R16000 SampleRate = 2
	R22050 SampleRate = 3
	R32000 SampleRate = 4
	R44056 SampleRate = 5
	R44100 SampleRate = 6

	B8  SampleResolution = 0
	B16 SampleResolution = 1
	B20 SampleResolution = 2
	B24 SampleResolution = 3
)

type Packet interface {
	Version() byte
	PacketType() PacketType
	RawPayload() []byte
	Raw() []byte
}

type simplePacket struct {
	version byte
	pType   PacketType
}

func (p *simplePacket) Version() byte          { return p.version }
func (p *simplePacket) PacketType() PacketType { return p.pType }
func (p *simplePacket) RawPayload() []byte     { return make([]byte, 0) }
func (p *simplePacket) Raw() []byte            { return []byte{p.version, byte(p.pType)} }

type DetailsPacket struct {
	version    byte
	pType      PacketType
	rate       SampleRate
	resolution SampleResolution
}

func (p *DetailsPacket) Version() byte                      { return p.version }
func (p *DetailsPacket) PacketType() PacketType             { return p.pType }
func (p *DetailsPacket) SampleRate() SampleRate             { return p.rate }
func (p *DetailsPacket) SampleResolution() SampleResolution { return p.resolution }
func (p *DetailsPacket) RawPayload() []byte {
	return []byte{
		byte(p.rate),
		byte(p.resolution),
	}
}
func (p *DetailsPacket) Raw() []byte {
	return []byte{
		byte(p.version),
		byte(p.pType),
		byte(p.rate),
		byte(p.resolution),
	}
}

type StreamPacket struct {
	version   byte
	pType     PacketType
	timestamp uint64
	data      []byte
}

func (p *StreamPacket) Version() byte          { return p.version }
func (p *StreamPacket) PacketType() PacketType { return p.pType }
func (p *StreamPacket) Timestamp() uint64      { return p.timestamp }
func (p *StreamPacket) StreamData() []byte     { return p.data }
func (p *StreamPacket) RawPayload() []byte {
	payload := make([]byte, len(p.data)+8)
	binary.LittleEndian.PutUint64(payload[0:8], p.timestamp)
	copy(payload[8:], p.data)
	return payload
}

func (p *StreamPacket) Raw() []byte {
	raw := make([]byte, len(p.data)+10)
	raw[0] = p.version
	raw[1] = byte(p.pType)
	binary.LittleEndian.PutUint64(raw[2:10], p.timestamp)
	copy(raw[10:], p.data)
	return raw
}

func NewSimplePacket(version byte, t PacketType) Packet {
	return &simplePacket{version, t}
}
func NewRequestPacket(version byte) Packet {
	return &simplePacket{version, REQUEST}
}

func NewAckPacket(version byte) Packet {
	return &simplePacket{version, ACK}
}

func NewClosePacket(version byte) Packet {
	return &simplePacket{version, CLOSE}
}

func NewPingPacket(version byte) Packet {
	return &simplePacket{version, PING}
}

func NewPongPacket(version byte) Packet {
	return &simplePacket{version, PONG}
}

func NewDetailsPacket(version byte, rate SampleRate, res SampleResolution) *DetailsPacket {
	return &DetailsPacket{version, DETAILS, rate, res}
}

func NewStreamPacket(version byte, timestamp uint64, data []byte) *StreamPacket {
	return &StreamPacket{version, STREAM, timestamp, data}
}
