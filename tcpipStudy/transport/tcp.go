package transport

import (
	"context"
	"log"

	"github.com/emahiro/il/tcpipStudy/inet"
	"github.com/emahiro/il/tcpipStudy/nw"
)

const (
	QueueSize = 10
)

type TcpPacket struct {
	IpHeader  *inet.Header
	TcpHeader *Header
	Packet    nw.Packet
}

type TcpPacketQueue struct {
	manager       *ConnectionManager
	outgoingQueue chan nw.Packet
	ctx           context.Context
	cancel        context.CancelFunc
}

func NewTcpPacketQueue() *TcpPacketQueue {
	ConnectionManager := NewConnectionManager()
	return &TcpPacketQueue{
		manager:       ConnectionManager,
		outgoingQueue: make(chan nw.Packet, QueueSize),
	}
}

func (q *TcpPacketQueue) ManageQueues(ip *inet.IpPacketQueue) {
	q.ctx, q.cancel = context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-q.ctx.Done():
				return
			default:
				ipp, err := ip.Read()
				if err != nil {
					log.Printf("read error: %v", err)
				}
				tcpHdr, err := unmarshal(ipp.Packet.Buf[ipp.IpHeader.IHL*4 : ipp.Packet.N])
				if err != nil {
					log.Printf("unmarshal error: %v", err)
					continue
				}
				tcpPkt := TcpPacket{
					IpHeader:  ipp.IpHeader,
					TcpHeader: tcpHdr,
					Packet:    ipp.Packet,
				}
				q.manager.recv(q, tcpPkt)
			}
		}

	}()

	// 受信
	go func() {
		for {
			select {
			case <-q.ctx.Done():
				return
			case pkt := <-q.outgoingQueue:
				if err := ip.Write(pkt); err != nil {
					log.Printf("write error: %v", err)
					return
				}
			}
		}
	}()
}
