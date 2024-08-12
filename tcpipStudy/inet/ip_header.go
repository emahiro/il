package inet

import (
	"encoding/binary"
	"fmt"
)

const (
	IP_VERSION        = 4       // IPv4 or IPv6
	IHL               = 5       // Internet Header Length
	TOC               = 0       // Type of Service 現在は DS ( Differentiated services) に置き換えられている
	TTL               = 64      // Time to Live
	LENGTH            = IHL * 4 // Header Length
	TCP_PROTOCOL      = 6
	IP_HEADER_MIN_LEN = 20
)

type Header struct {
	Version        uint8
	IHL            uint8
	TOS            uint8
	TotalLength    uint16
	ID             uint16
	Flags          uint8
	FragmentOffset uint16
	TTL            uint8
	Protocol       uint8
	Checksum       uint16
	SrcIP          [4]byte
	DstIP          [4]byte
}

func NewIP(srcIP, dstIP [4]byte, len int) *Header {
	return &Header{
		Version:     IP_VERSION,
		IHL:         IHL,
		TOS:         TOC,
		TotalLength: uint16(LENGTH + len),
		ID:          0,
		Flags:       0x40,
		TTL:         TTL,
		Protocol:    TCP_PROTOCOL,
		Checksum:    0,
		SrcIP:       srcIP,
		DstIP:       dstIP,
	}
}

func unmarshal(pkt []byte) (*Header, error) {
	if len(pkt) < IP_HEADER_MIN_LEN {
		return nil, fmt.Errorf("invalid ip header length: %d", len(pkt))
	}

	// parse header along with https://www.rfc-editor.org/rfc/rfc791
	h := &Header{
		Version:        pkt[0] >> 4,
		IHL:            pkt[0] & 0x0f,
		TOS:            pkt[1],
		TotalLength:    binary.BigEndian.Uint16([]byte(pkt[2:4])),
		ID:             binary.BigEndian.Uint16([]byte(pkt[4:6])),
		Flags:          pkt[6] >> 5,
		FragmentOffset: binary.BigEndian.Uint16(pkt[6:8]) & 0x1fff,
		TTL:            pkt[8],
		Protocol:       pkt[9],
		Checksum:       binary.BigEndian.Uint16(pkt[10:12]),
	}
	copy(h.SrcIP[:], pkt[12:16])
	copy(h.DstIP[:], pkt[16:20])

	return h, nil
}

func (h *Header) Marshal() []byte {
	versionAndIHL := (IP_VERSION << 4) | h.IHL
	flagsAndFragmentOffset := (uint16(h.FragmentOffset) << 13) | (h.FragmentOffset & 0x1fff)

	pkt := make([]byte, 20)
	pkt[0] = versionAndIHL
	pkt[1] = 0
	binary.BigEndian.PutUint16(pkt[2:4], h.TotalLength)
	binary.BigEndian.PutUint16(pkt[4:6], h.ID)
	binary.BigEndian.PutUint16(pkt[6:8], flagsAndFragmentOffset)
	pkt[8] = h.TTL
	pkt[9] = h.Protocol
	binary.BigEndian.PutUint16(pkt[10:12], h.Checksum)
	copy(pkt[12:16], h.SrcIP[:])
	copy(pkt[16:20], h.DstIP[:])

	h.setChecksum(pkt)
	binary.BigEndian.PutUint16(pkt[10:12], h.Checksum)

	return pkt
}

func (h *Header) setChecksum(pkt []byte) {
	length := len(pkt)
	var checkSum uint32

	for i := 0; i < length; i += 2 {
		checkSum += uint32(binary.BigEndian.Uint16(pkt[i : i+2]))
	}

	for checkSum > 0xffff {
		checkSum = (checkSum & 0xffff) + (checkSum >> 16)
	}

	h.Checksum = ^uint16(checkSum)
}
