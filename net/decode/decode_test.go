package decode

import (
	"bytes"
	"github.com/rickbutton/speaker/net/packet"
	"testing"
)

func TestDecodeSimplePacket(t *testing.T) {
	types := []packet.PacketType{packet.REQUEST,
		packet.ACK,
		packet.CLOSE,
		packet.PING,
		packet.PONG}
	for i := range types {
		raw := []byte{123, byte(types[i])}
		p := DecodeRawPacket(raw)
		if p.Version() != 123 {
			t.Errorf("Version() = %d, want %d", p.Version(), 123)
		}
		if p.PacketType() != types[i] {
			t.Errorf("PacketType() = %d, want %d", p.PacketType(), types[i])
		}
	}
}

func TestDecodeDetailsPacket(t *testing.T) {
	raw := []byte{123, byte(packet.DETAILS), byte(packet.R8000), byte(packet.B16)}
	p := DecodeRawPacket(raw).(*packet.DetailsPacket)
	if p.Version() != 123 {
		t.Errorf("Version() = %d, want %d", p.Version(), 123)
	}
	if p.PacketType() != packet.DETAILS {
		t.Errorf("PacketType() = %d, want %d", p.PacketType(), packet.DETAILS)
	}
	if p.SampleRate() != packet.R8000 {
		t.Errorf("SampleRate() = %d, want %d", p.SampleRate(), packet.R8000)
	}
	if p.SampleResolution() != packet.B16 {
		t.Errorf("SampleResolution() = %d, want %d", p.SampleResolution(), packet.B16)
	}
}

func TestDecodeStreamPacket(t *testing.T) {
	raw := []byte{123, byte(packet.STREAM), 63, 66, 15, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5}
	p := DecodeRawPacket(raw).(*packet.StreamPacket)
	if p.PacketType() != packet.STREAM {
		t.Errorf("PacketType() = %d, want %d", p.PacketType(), packet.STREAM)
	}
	if p.Version() != 123 {
		t.Errorf("%d Version() = %d, want %d", p.PacketType(), p.Version(), 123)
	}
	if p.Timestamp() != 999999 {
		t.Errorf("Timestamp() = %d, want %d", p.Timestamp(), 999999)
	}
	if !bytes.Equal(p.StreamData(), raw[10:15]) {
		t.Errorf("StreamData() = %v, want %v", p.StreamData(), raw[10:15])
	}
	payload := []byte{63, 66, 15, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5}
	if !bytes.Equal(p.RawPayload(), payload) {
		t.Errorf("RawPayload() = %v, want %v", p.RawPayload(), payload)
	}
	if !bytes.Equal(p.Raw(), raw) {
		t.Errorf("Raw() = %v, want %v", p.Raw(), raw)
	}
}
