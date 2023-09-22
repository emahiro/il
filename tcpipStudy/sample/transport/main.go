package main

import (
	"fmt"

	"github.com/emahiro/il/tcpipStudy/inet"
	"github.com/emahiro/il/tcpipStudy/nw"
	"github.com/emahiro/il/tcpipStudy/transport"
)

func main() {
	nw, _ := nw.NewTun()
	nw.Bind()

	ip := inet.NewIpPacketQueue()
	ip.ManageQueues(nw)

	tcp := transport.NewTcpPacketQueue()
	tcp.ManageQueues(ip)

	for {
		pkt, _ := tcp.ReadAcceptConnection()
		fmt.Printf("TCH Header: %+v\n", pkt.Pkt.TcpHeader)
	}
}
