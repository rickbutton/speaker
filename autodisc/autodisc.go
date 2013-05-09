package autodisc

import (
  "net"
  "fmt"
  "errors"
  "bytes"
  "time"
)

const (
  DiscoveryPort = 33333
  TalkPort      = 33334
)

var (
  broadcast = net.IPv4(255, 255, 255, 255)
  listen    = net.IPv4(0,0,0,0)
  ping = []byte{0xDE, 0xAD}
  reply = []byte{0xC0, 0xDE}
)

func FindServer() (net.IP, error) {
  bd := net.UDPAddr{IP: broadcast, Port:DiscoveryPort}
  ip := net.IP(nil)
  socket, err := net.DialUDP("udp", nil, &bd);
  if (err != nil) { return nil, err }
  socket.Write(ping)

  buf := make([]byte, 2)
  n, addr, err := socket.ReadFromUDP(buf)
  if (err != nil) { return nil, err }
  if (!checkDisc(n, buf, reply)) {
    ip, err = net.IP(nil), errors.New(fmt.Sprintf("invalid reply magic n==%d m==%x%x", n, buf[0], buf[1]))
  } else {
    ip, err = addr.IP, error(nil)
  }
  return ip, err
}

type AutodiscServer struct {
  clients map[string] time.Time
}
func (*AutodiscServer) Listen() {
  bd := net.UDPAddr{IP: listen, Port: DiscoveryPort}
  socket, err := net.ListenUDP("udp4", &bd)
  if (err != nil) { panic(err) }
  for {
    buf := make([]byte, 2)
    n, addr, err := socket.ReadFromUDP(buf)
    if (err != nil) { panic(err) }
    if (checkDisc(n, buf, ping)) {
      if (err != nil) { panic(err) }
      socket.WriteTo(reply, addr)
    }
  }
}


func checkDisc(n int, buf []byte, magic []byte) bool {
  return n == 2 && bytes.Equal(buf, magic)
}
