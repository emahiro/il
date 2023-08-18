package inet

import nw "github.com/emahiro/il/tcpipStudy/nw"

const (
	QueueSize = 10
)

type IpPacket struct {
	IpHeader *Header
	Packet   nw.Packet
}
