package main

import (
	"encoding/hex"
	"fmt"

	"github.com/emahiro/il/tcpipStudy/nw"
)

func main() {
	nw, _ := nw.NewTun()
	nw.Bind()

	for {
		pkt, err := nw.Read()
		if err != nil {
			panic(err)
		}
		fmt.Print(hex.Dump(pkt.Buf[:pkt.N]))
		_ = nw.Write(pkt)
	}
}
