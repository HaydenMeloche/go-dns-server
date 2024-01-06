package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2053")
	if err != nil {
		log.Fatal("Failed to resolve UDP address:", err)
		return
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal("Failed to bind to address:", err)
		return
	}
	defer func(udpConn *net.UDPConn) {
		err := udpConn.Close()
		if err != nil {
			log.Fatal("Failed to close UDP connection:", err)
		}
	}(udpConn)

	buf := make([]byte, 512)

	for {
		size, source, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal("Error receiving data:", err)
			break
		}

		receivedData := buf[:size]
		log.Printf("Received %d bytes from %s: %s\n", size, source, receivedData)

		header := DeserializeDNSHeader(receivedData[:12])
		header.FLAGS.QR = true // Set the QR flag to indicate that this is a response
		header.FLAGS.RCODE = 4 // Set the response code to 4 (Not Implemented)

		log.Println("Received header:", receivedData[12:])
		question := DeserializeQuestions(receivedData[12:])

		answers := make([]Answer, len(question))
		for i, q := range question {
			answers[i] = Answer{
				NAME:     q.NAME,
				TYPE:     1,
				CLASS:    1,
				TTL:      60,
				RDLENGTH: 4,
				RDATA:    []byte{8, 8, 8, 8}, // the IP address here is hardcoded to the Google DNS server
			}
			header.ANCOUNT++
		}

		fmt.Println(">>> Sending response:", header, question, answers)
		response := header.Serialize()
		for i := range question {
			response = append(response, question[i].Serialize()...)
		}
		for i := range answers {
			response = append(response, answers[i].Serialize()...)
		}

		_, err = udpConn.WriteToUDP(response, source)
		if err != nil {
			log.Fatal("Failed to send response:", err)
		}
	}
}
