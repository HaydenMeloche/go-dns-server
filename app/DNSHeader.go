package main

import "encoding/binary"

type DNSHeader struct {
	// Total 12 bytes (96 bits)
	ID      uint16 // Packet Identifier (16 bit)
	FLAGS   Flags  // Flags (16 bit)
	QDCOUNT uint16 // Question Count (16 bit)
	ANCOUNT uint16 // Answer Record Count (16 bit)
	NSCOUNT uint16 // Authority Record Count (16 bit)
	ARCOUNT uint16 // Additional Record Count (16 bit)

}

func (dh DNSHeader) Serialize() []byte {
	header := make([]byte, 12)
	binary.BigEndian.PutUint16(header[0:2], dh.ID)
	binary.BigEndian.PutUint16(header[2:4], dh.FLAGS.Serialize())
	binary.BigEndian.PutUint16(header[4:6], dh.QDCOUNT)
	binary.BigEndian.PutUint16(header[6:8], dh.ANCOUNT)
	binary.BigEndian.PutUint16(header[8:10], dh.NSCOUNT)
	binary.BigEndian.PutUint16(header[10:12], dh.ARCOUNT)
	return header
}

func DeserializeDNSHeader(data []byte) DNSHeader {
	return DNSHeader{
		ID:      binary.BigEndian.Uint16(data[0:2]),
		FLAGS:   DeserializeFlags(data[2:4]),
		QDCOUNT: binary.BigEndian.Uint16(data[4:6]),
		ANCOUNT: binary.BigEndian.Uint16(data[6:8]),
		NSCOUNT: binary.BigEndian.Uint16(data[8:10]),
		ARCOUNT: binary.BigEndian.Uint16(data[10:12]),
	}
}
