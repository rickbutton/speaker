package main

import (
  "github.com/rickbutton/speaker/autodisc"
  "fmt"
  "net"
)

func main() {
  server, err := autodisc.FindServer()
  if (err != nil) { panic(err) }
  addr := net.IPAddr{server}
  fmt.Printf("Found server %s", addr.String())
  for{}
}
