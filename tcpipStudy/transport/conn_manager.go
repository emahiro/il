package transport

import "sync"

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
	Connections []Connection
	AcceptQueue chan *Connection
	lock        sync.Mutex
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		AcceptQueue: make(chan *Connection, QueueSize),
	}
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

}
