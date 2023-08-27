package transport

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	StateListen State = iota
	StateSynReceived
	StateEstablished
	StateClosedWait
	StateLastAck
	StateClosed
)

type State int // TCP の接続状態

type Connection struct {
	SrcPort uint16
	DstPort uint16
	State   State
	Pkt     TcpPacket
	N       uintptr

	initialSeqNum   uint32
	incrementSeqNum uint32

	isAccept bool
}

type ConnectionManager struct {
	Connections           []Connection
	AcceptConnectionQueue chan Connection
	lock                  sync.Mutex
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		AcceptConnectionQueue: make(chan Connection, QueueSize),
	}
}

func (m *ConnectionManager) addConnection(pkt TcpPacket) Connection {
	m.lock.Lock()
	defer m.lock.Unlock()

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	conn := Connection{
		SrcPort:         pkt.TcpHeader.SrcPort,
		DstPort:         pkt.TcpHeader.DstPort,
		State:           StateSynReceived,
		N:               pkt.Packet.N,
		Pkt:             pkt,
		initialSeqNum:   uint32(r.Int31()),
		incrementSeqNum: 0,
	}
	m.Connections = append(m.Connections, conn)

	return conn
}

func (m *ConnectionManager) find(pkt TcpPacket) (Connection, bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	for _, conn := range m.Connections {
		if conn.SrcPort == pkt.TcpHeader.SrcPort && conn.DstPort == pkt.TcpHeader.DstPort {
			// 既に確立してるコネクションの場合はそのコネクションを返す。
			return conn, true
		}
	}
	return Connection{}, false
}

func (m *ConnectionManager) remove(pkt TcpPacket) (Connection, bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	for i, conn := range m.Connections {
		if conn.SrcPort == pkt.TcpHeader.SrcPort && conn.DstPort == pkt.TcpHeader.DstPort {
			m.Connections = append(m.Connections[:i], m.Connections[i+1:]...)
		}
	}
	return Connection{}, false
}

func (m *ConnectionManager) update(pkt TcpPacket, state State, isAccept bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	for i, conn := range m.Connections {
		if conn.SrcPort == pkt.TcpHeader.SrcPort && conn.DstPort == pkt.TcpHeader.DstPort {
			m.Connections[i].State = state
			m.Connections[i].isAccept = isAccept
			return
		}
	}
}

func (m *ConnectionManager) updateIncrementSeqNum(pkt TcpPacket, val uint32) {
	m.lock.Lock()
	defer m.lock.Unlock()

	for i, conn := range m.Connections {
		if conn.SrcPort == pkt.TcpHeader.SrcPort && conn.DstPort == pkt.TcpHeader.DstPort {
			m.Connections[i].incrementSeqNum = val
			return
		}
	}
}

func (m *ConnectionManager) recv(q *TcpPacketQueue, pkt TcpPacket) {
	conn, connected := m.find(pkt)
	if !connected {
		m.addConnection(pkt)
	} else {
		conn.Pkt = pkt
	}

	if pkt.TcpHeader.Flags.SYN && connected {
		log.Println("Received SYN packet")
		// queue write
		m.update(pkt, StateSynReceived, false)
	}

	if connected && pkt.TcpHeader.Flags.ACK && conn.State == StateSynReceived {
		log.Println("Received ACK packet")
		m.update(pkt, StateEstablished, false)
	}

	if connected && pkt.TcpHeader.Flags.PSH && conn.State == StateEstablished {
		log.Println("Received PSH packet")
		q.Write(conn, HeaderFlags{
			ACK: true,
		}, nil)
		m.update(pkt, StateEstablished, true)
		m.AcceptConnectionQueue <- conn
	}

	if connected && pkt.TcpHeader.Flags.FIN && conn.State == StateEstablished {
		log.Println("Received FIN packet")

		q.Write(conn, HeaderFlags{
			ACK: true,
		}, nil)
		m.update(pkt, StateClosedWait, false)

		q.Write(conn, HeaderFlags{
			FIN: true,
			ACK: true,
		}, nil)

		m.update(pkt, StateLastAck, false)
	}

	if connected && pkt.TcpHeader.Flags.ACK && conn.State == StateLastAck {
		log.Println("Received ACK packet")
		m.update(pkt, StateClosed, false)
		m.remove(pkt)
	}
}
