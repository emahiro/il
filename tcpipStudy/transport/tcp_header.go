package transport

import (
	"encoding/binary"
	"errors"

	"github.com/emahiro/il/tcpipStudy/inet"
)

const (
	Length      = 20
	WindowSize  = 65535
	ProtocolVer = 6
)

type Header struct {
	SrcPort    uint16 // 送信元 Port
	DstPort    uint16 // 送信先 Port
	SeqNum     uint32 // シーケンス番号
	AckNum     uint32 // 応答確認番号
	DataOff    uint8
	Reserved   uint8
	Flags      HeaderFlags
	WindowSize uint16 // 受信側が受け入れ可能なバイト数
	Checksum   uint16
	UrgentPtr  uint16
}

type HeaderFlags struct {
	CWR bool
	ECE bool
	URG bool
	ACK bool // Acknowledgment 応答確認
	PSH bool // Push
	RST bool // Reset
	SYN bool // Synchronize 同期
	FIN bool // Finish
}

func NewTCPHeader(srcPort, dstPort uint16, seqNum, ackNum uint32, flags HeaderFlags) *Header {
	dataOff := uint16(Length / 4)
	dataOff <<= 4

	return &Header{
		SrcPort:    srcPort,
		DstPort:    dstPort,
		SeqNum:     seqNum,
		AckNum:     ackNum,
		DataOff:    uint8(dataOff),
		Reserved:   0x12,
		Flags:      flags,
		WindowSize: uint16(WindowSize),
		Checksum:   0,
		UrgentPtr:  0,
	}
}

func unmarshal(pkt []byte) (*Header, error) {
	if len(pkt) < 20 {
		return nil, errors.New("invalid TCP header length for too short")
	}

	flags := unmarshalFlags(pkt[13])

	// described at https://www.rfc-editor.org/rfc/rfc9293
	return &Header{
		SrcPort:    binary.BigEndian.Uint16(pkt[0:2]),
		DstPort:    binary.BigEndian.Uint16(pkt[2:4]),
		SeqNum:     binary.BigEndian.Uint32(pkt[4:8]),
		AckNum:     binary.BigEndian.Uint32(pkt[8:12]),
		DataOff:    pkt[12] >> 4,
		Flags:      flags,
		WindowSize: binary.BigEndian.Uint16(pkt[14:16]),
		Checksum:   binary.BigEndian.Uint16(pkt[16:18]),
		UrgentPtr:  binary.BigEndian.Uint16(pkt[18:20]),
	}, nil
}

func unmarshalFlags(flags uint8) HeaderFlags {
	return HeaderFlags{
		CWR: flags&0x80 == 0x80,
		ECE: flags&0x40 == 0x40,
		URG: flags&0x20 == 0x20,
		ACK: flags&0x10 == 0x10,
		PSH: flags&0x08 == 0x08,
		RST: flags&0x04 == 0x04,
		SYN: flags&0x02 == 0x02,
		FIN: flags&0x01 == 0x01,
	}
}

func (h *HeaderFlags) marshal() uint8 {
	var packedFlags uint8
	if h.CWR {
		packedFlags |= 0x80 // |= ORビット演算の複合代入演算子 => ORを取る
	}
	if h.ECE {
		packedFlags |= 0x40
	}
	if h.URG {
		packedFlags |= 0x20
	}
	if h.ACK {
		packedFlags |= 0x10
	}
	if h.PSH {
		packedFlags |= 0x08
	}
	if h.RST {
		packedFlags |= 0x04
	}
	if h.SYN {
		packedFlags |= 0x02
	}
	if h.FIN {
		packedFlags |= 0x01
	}
	return packedFlags
}

func (h *Header) setChkSum(ipHdr *inet.Header, pkt []byte) {
	pseudoHdr := make([]byte, 12)
	copy(pseudoHdr[0:4], ipHdr.SrcIP[:])
	copy(pseudoHdr[4:8], ipHdr.DstIP[:])
	pseudoHdr[8] = 0
	pseudoHdr[9] = ipHdr.Protocol
	binary.BigEndian.PutUint16(pseudoHdr[10:12], uint16(len(pkt)))

	buf := append(pseudoHdr, pkt...)
	if len(buf)%2 != 0 {
		buf = append(buf, 0)
	}

	var sum uint32
	for i := 0; i < len(buf); i += 2 {
		sum += uint32(binary.BigEndian.Uint16(buf[i : i+2]))
	}

	if sum > 0xffff {
		sum = (sum >> 16) + (sum & 0xffff)
	}

	h.Checksum = ^uint16(sum)
}

func (h *Header) Marshal(ipHdr *inet.Header, data []byte) []byte {
	flags := h.Flags.marshal()

	pkt := make([]byte, 20)
	binary.BigEndian.PutUint16(pkt[0:2], h.SrcPort)
	binary.BigEndian.PutUint16(pkt[2:4], h.DstPort)
	binary.BigEndian.PutUint32(pkt[4:8], h.SeqNum)
	binary.BigEndian.PutUint32(pkt[8:12], h.AckNum)
	pkt[12] = h.DataOff
	pkt[13] = flags
	binary.BigEndian.PutUint16(pkt[14:16], h.WindowSize)
	binary.BigEndian.PutUint16(pkt[16:18], h.Checksum)
	binary.BigEndian.PutUint16(pkt[18:20], h.UrgentPtr)

	h.setChkSum(ipHdr, append(pkt, data...))
	binary.BigEndian.PutUint16(pkt[16:18], h.Checksum)

	return pkt
}
