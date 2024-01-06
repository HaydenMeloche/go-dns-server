package main

import (
	"encoding/binary"
	"strings"
)

type Question struct {
	NAME  string // Domain Name represented as a sequence of labels
	TYPE  uint16 // Record Type (16 bit) see all types here: https://www.rfc-editor.org/rfc/rfc1035#section-3.2.2
	CLASS uint16 // Record Class (16 bit) usually set as '1'
}

func (q Question) Serialize() []byte {
	buf := make([]byte, 0)
	for _, label := range strings.Split(q.NAME, ".") {
		buf = append(buf, byte(len(label)))
		buf = append(buf, label...)
	}
	buf = append(buf, 0x0)
	buf = binary.BigEndian.AppendUint16(buf, q.TYPE)
	buf = binary.BigEndian.AppendUint16(buf, q.CLASS)

	return buf
}

func DeserializeQuestions(data []byte) []Question {
	var questions []Question
	offset := 0

	for offset < len(data) {
		name, nameLength := parseName(data, offset)
		offset += nameLength

		questions = append(questions, Question{
			NAME:  name,
			TYPE:  binary.BigEndian.Uint16(data[offset : offset+2]),
			CLASS: binary.BigEndian.Uint16(data[offset+2 : offset+4]),
		})

		offset += 4 // Move past the TYPE and CLASS fields
	}

	return questions
}

func parseName(data []byte, offset int) (string, int) {
	var labels []string
	originalOffset := offset

	for {
		if (data[offset] & 0xC0) == 0xC0 {
			// This is a pointer to an existing name.
			pointer := (int(binary.BigEndian.Uint16(data[offset:offset+2])) & 0x3FFF) / 4
			label, _ := parseName(data, pointer)
			labels = append(labels, label)
			offset += 2
			break
		} else {
			// This is a label.
			length := int(data[offset])
			if length == 0 || length > 63 {
				// End of this name.
				offset++
				break
			}
			label := string(data[offset+1 : offset+1+length])
			labels = append(labels, label)
			offset += 1 + length
		}
	}

	return strings.Join(labels, "."), offset - originalOffset
}
