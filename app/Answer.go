package main

import (
	"encoding/binary"
	"strings"
)

type Answer struct {
	NAME     string // Domain Name represented as a sequence of labels
	TYPE     uint16 // Record Type (16 bit) see all types here: https://www.rfc-editor.org/rfc/rfc1035#section-3.2.2
	CLASS    uint16 // Record Class (16 bit) usually set as '1'
	TTL      uint32 // Time to Live (32 bit)
	RDLENGTH uint16 // RDATA Length (16 bit)
	RDATA    []byte // RDATA (variable length)
}

func (a Answer) Serialize() []byte {
	buf := make([]byte, 0)
	for _, label := range strings.Split(a.NAME, ".") {
		buf = append(buf, byte(len(label)))
		buf = append(buf, label...)
	}
	buf = append(buf, 0x0)
	buf = binary.BigEndian.AppendUint16(buf, a.TYPE)
	buf = binary.BigEndian.AppendUint16(buf, a.CLASS)

	buf = binary.BigEndian.AppendUint32(buf, a.TTL)
	buf = binary.BigEndian.AppendUint16(buf, a.RDLENGTH)
	buf = append(buf, a.RDATA...)

	return buf
}

func DeserializeAnswer(data []byte) Answer {
	var a Answer
	var name []string
	for i := 0; i < len(data); i++ {
		if data[i] == 0x0 {
			a.NAME = strings.Join(name, ".")
			a.TYPE = binary.BigEndian.Uint16(data[i+1 : i+3])
			a.CLASS = binary.BigEndian.Uint16(data[i+3 : i+5])
			a.TTL = binary.BigEndian.Uint32(data[i+5 : i+9])
			a.RDLENGTH = binary.BigEndian.Uint16(data[i+9 : i+11])
			a.RDATA = data[i+11 : i+11+int(a.RDLENGTH)]
			break
		}
		name = append(name, string(data[i+1:i+1+int(data[i])]))
		i += int(data[i])
	}
	return a
}
