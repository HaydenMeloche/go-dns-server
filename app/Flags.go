package main

type Flags struct {
	QR     bool  // Query/Response Indicator (1 bit)
	OPCODE uint8 // Query/Response Indicator (4 bit)
	AA     bool  // Authoritative Answer (1 bit)
	TC     bool  // Truncation (1 bit)
	RD     bool  // Recursion Desired (1 bit)
	RA     bool  // Recursion Desired (1 bit)
	Z      uint8 // Reserved (3 bit)
	RCODE  uint8 // Response Code (4 bit)
}

func (f Flags) Serialize() uint16 {
	var header uint16
	if f.QR {
		header |= 0x1 << 15
	}
	opcode := uint16(f.OPCODE) & 0b1111
	header |= opcode << 11

	if f.AA {
		header |= 0x1 << 10
	}
	if f.TC {
		header |= 0x1 << 9
	}
	if f.RD {
		header |= 0x1 << 8
	}
	if f.RA {
		header |= 0x1 << 7
	}
	zFlag := uint16(f.Z) & 0b111
	header |= zFlag << 4
	rcodeFlag := uint16(f.RCODE) & 0b1111
	header |= rcodeFlag

	return header
}

func DeserializeFlags(bytes []byte) Flags {
	return Flags{
		QR:     bytes[0]&0b10000000 > 0,
		OPCODE: bytes[0] & 0b01111000 >> 3,
		AA:     bytes[0]&0b00000100 > 0,
		TC:     bytes[0]&0b00000010 > 0,
		RD:     bytes[0]&0b00000001 > 0,
		RA:     bytes[1]&0b10000000 > 0,
		Z:      bytes[1] & 0b01110000 >> 4,
		RCODE:  bytes[1] & 0b00001111,
	}
}
