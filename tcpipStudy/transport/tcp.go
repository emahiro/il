package transport

import (
	"context"
	"fmt"
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

func (q *TcpPacketQueue) Close() {
	q.cancel()
}

func (q *TcpPacketQueue) Write(conn Connection, flags HeaderFlags, data []byte) {
	pkt := conn.Pkt
	tcpDataLen := int(pkt.Packet.N) - (int(pkt.IpHeader.IHL) * 4) - (int(pkt.TcpHeader.DataOff) * 4)

	incrementAckNum := 0
	if tcpDataLen == 0 {
		incrementAckNum = 1
	} else {
		incrementAckNum = tcpDataLen
	}

	ackNum := pkt.TcpHeader.SeqNum + uint32(incrementAckNum)
	seqNum := conn.initialSeqNum + conn.incrementSeqNum

	writeIpHdr := inet.NewIP(pkt.IpHeader.DstIP, pkt.IpHeader.SrcIP, Length+len(data))
	writeTcpHdr := NewTCPHeader(
		pkt.TcpHeader.DstPort,
		pkt.TcpHeader.SrcPort,
		seqNum,
		ackNum,
		flags,
	)

	ipHdr := writeIpHdr.Marshal()
	tcpHdr := writeTcpHdr.Marshal(conn.Pkt.IpHeader, data)

	writePkt := append(ipHdr, tcpHdr...)
	writePkt = append(writePkt, data...)

	incrementSeqNum := 0
	if flags.SYN || flags.FIN {
		incrementSeqNum += 1
	}
	incrementSeqNum += len(data)
	q.manager.updateIncrementSeqNum(pkt, uint32(incrementSeqNum))
	q.outgoingQueue <- nw.Packet{
		Buf: ipHdr,
		N:   uintptr(len(writePkt)),
	}
}

func (q *TcpPacketQueue) ReadAcceptConnection() (Connection, error) {
	pkt, received := <-q.manager.AcceptConnectionQueue
	if !received {
		return Connection{}, fmt.Errorf("accept connection queue is closed")
	}
	return pkt, nil
}
