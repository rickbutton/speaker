package main

import (
  "github.com/rickbutton/speaker/autodisc"
)

func main() {
  adServer := new(autodisc.AutodiscServer)
  adServer.Listen()
}
