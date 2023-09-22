package main

import (
	"log"

	"github.com/emahiro/il/tcpipStudy/inet"
	"github.com/emahiro/il/tcpipStudy/nw"
)

func main() {
	nw, err := nw.NewTun()
	if err != nil {
		panic(err)
	}
	nw.Bind()
	ip := inet.NewIpPacketQueue()
	ip.ManageQueues(nw)

	for {
		pkt, err := ip.Read()
		if err != nil {
			panic(err)
		}
		log.Printf("IP Header: %+v\n", pkt.IpHeader)
	}
}
