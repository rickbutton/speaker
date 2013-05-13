package packet

import (
	"bytes"
	"testing"
)

func TestNewSimplePacket(t *testing.T) {
  types := []PacketType{
    REQUEST,
    ACK,
    CLOSE,
    PING,
    PONG,
  }
  for i := range types {
    p := NewSimplePacket(123, types[i])
    if p.PacketType() != types[i] {
      t.Errorf("PacketType() = %d, want %d", p.PacketType(), types[i])
    }
    if p.Version() != 123 {
      t.Errorf("%d Version() = %d, want %d", p.PacketType(), p.Version(), 123)
    }
    if len(p.RawPayload()) != 0 {
      t.Errorf("Packet type %d shouldn't have a payload, but does")
    }
    raw := []byte{123, byte(types[i])}
    if !bytes.Equal(p.Raw(), raw) {
      t.Errorf("Raw() = %v, want %v", p.Raw(), raw)
    }
  }
}
func TestAllNewSimplePacket(t *testing.T) {
	types := map[PacketType]func(byte) Packet{
		REQUEST: NewRequestPacket,
		ACK:     NewAckPacket,
		CLOSE:   NewClosePacket,
		PING:    NewPingPacket,
		PONG:    NewPongPacket,
	}
	for ty, f := range types {
		p := f(123)
		if p.PacketType() != ty {
			t.Errorf("PacketType() = %d, want %d", p.PacketType(), ty)
		}
		if p.Version() != 123 {
			t.Errorf("%d Version() = %d, want %d", p.PacketType(), p.Version(), 123)
		}
		if len(p.RawPayload()) != 0 {
			t.Errorf("Packet type %d shouldn't have a payload, but does")
		}
		raw := []byte{123, byte(ty)}
		if !bytes.Equal(p.Raw(), raw) {
			t.Errorf("Raw() = %v, want %v", p.Raw(), raw)
		}
	}
}

func TestNewDetailsPacket(t *testing.T) {
	p := NewDetailsPacket(123, R16000, B16)
	if p.PacketType() != DETAILS {
		t.Errorf("PacketType() = %d, want %d", p.PacketType(), DETAILS)
	}
	if p.Version() != 123 {
		t.Errorf("%d Version() = %d, want %d", p.PacketType(), p.Version(), 123)
	}
	if p.SampleRate() != R16000 {
		t.Errorf("SampleRate() = %d, want %d", p.SampleRate(), R16000)
	}
	if p.SampleResolution() != B16 {
		t.Errorf("SampleResolution() = %d, want %d", p.SampleResolution(), B16)
	}
	payload := []byte{byte(R16000), byte(B16)}
	if !bytes.Equal(p.RawPayload(), payload) {
		t.Errorf("RawPayload() = %v, want %v", p.RawPayload(), payload)
	}
	raw := []byte{123, byte(DETAILS), byte(R16000), byte(B16)}
	if !bytes.Equal(p.Raw(), raw) {
		t.Errorf("Raw() = %v, want %v", p.Raw(), raw)
	}
}

func TestNewStreamPacket(t *testing.T) {
	data := []byte{1, 2, 3, 4, 5}
	p := NewStreamPacket(123, 999999, data)
	if p.PacketType() != STREAM {
		t.Errorf("PacketType() = %d, want %d", p.PacketType(), STREAM)
	}
	if p.Version() != 123 {
		t.Errorf("%d Version() = %d, want %d", p.PacketType(), p.Version(), 123)
	}
	if p.Timestamp() != 999999 {
		t.Errorf("Timestamp() = %d, want %d", p.Timestamp(), 999999)
	}
	if !bytes.Equal(p.StreamData(), data) {
		t.Errorf("StreamData() = %v, want %v", p.StreamData(), data)
	}
	payload := []byte{63, 66, 15, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5}
	if !bytes.Equal(p.RawPayload(), payload) {
		t.Errorf("RawPayload() = %v, want %v", p.RawPayload(), payload)
	}
	raw := []byte{123, byte(STREAM), 63, 66, 15, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5}
	if !bytes.Equal(p.Raw(), raw) {
		t.Errorf("Raw() = %v, want %v", p.Raw(), raw)
	}
}
