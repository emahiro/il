package inet

import (
	"context"
	"fmt"
	"log"

	nw "github.com/emahiro/il/tcpipStudy/nw"
)

const (
	QueueSize = 10
)

type IpPacket struct {
	IpHeader *Header
	Packet   nw.Packet
}

type IpPacketQueue struct {
	incomingQueue chan IpPacket
	outgoingQueue chan nw.Packet
	ctx           context.Context
	cancel        context.CancelFunc
}

func NewIpPacketQueue() *IpPacketQueue {
	return &IpPacketQueue{
		incomingQueue: make(chan IpPacket, QueueSize),
		outgoingQueue: make(chan nw.Packet, QueueSize),
	}
}

func (ip *IpPacketQueue) ManageQueues(nw *nw.NetDevice) {
	ip.ctx, ip.cancel = context.WithCancel(context.Background())

	// 送信 Queue
	go func() {
		for {
			select {
			case <-ip.ctx.Done():
				return
			default:
				pkt, err := nw.Read()
				if err != nil {
					log.Printf("read error: %v", err)
				}
				h, err := unmarshal(pkt.Buf[:pkt.N])
				if err != nil {
					log.Printf("unmarshal error: %v", err)
				}
				ipp := IpPacket{
					IpHeader: h,
					Packet:   pkt,
				}
				// 送信
				ip.incomingQueue <- ipp
			}
		}
	}()

	// 受信 Queue
	go func() {
		for {
			select {
			case <-ip.ctx.Done():
				return
			case pkt := <-ip.outgoingQueue:
				if err := nw.Write(pkt); err != nil {
					log.Printf("write error: %v", err)
				}
			}
		}
	}()
}

func (q *IpPacketQueue) Close() {
	q.cancel()
}

func (q *IpPacketQueue) Read() (IpPacket, error) {
	pkt, ok := <-q.incomingQueue
	if !ok {
		return IpPacket{}, fmt.Errorf("incoming queue is closed")
	}
	return pkt, nil
}

func (q *IpPacketQueue) Write(pkt nw.Packet) error {
	select {
	case q.outgoingQueue <- pkt:
		return nil
	case <-q.ctx.Done():
		return fmt.Errorf("network closed")
	}
}
