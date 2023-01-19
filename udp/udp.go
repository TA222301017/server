package udp

import (
	"log"
	"net"

	"server/udp/setup"
	"server/udp/template"
	"server/udp/usecases"
	"server/udp/utils"
)

type Request struct {
	Packet     *template.BasePacket
	RemoteAddr *net.UDPAddr
}

func handlePacket(conn *net.UDPConn, request <-chan *Request) {
	for r := range request {
		switch r.Packet.OpCode {
		case template.KeyExchange:
			log.Printf("| 0x%x | KEY EXCHANGE\n", r.Packet.OpCode)
			res, err := usecases.KeyExchange(*r.Packet, r.RemoteAddr)
			if err != nil {
				log.Printf("| %s\n", err)
			}
			conn.WriteToUDP(res.Bytes(), r.RemoteAddr)
		case template.AddAccessRule:
			log.Printf("| UNIMPLEMENTED!\n")
		case template.EditAccessRule:
			log.Printf("| UNIMPLEMENTED!\n")
		case template.DeleteAccessRule:
			log.Printf("| UNIMPLEMENTED!\n")
		case template.LogAccessEvent:
			log.Printf("| UNIMPLEMENTED!\n")
		case template.LogRSSIEvent:
			log.Printf("| UNIMPLEMENTED!\n")
		case template.LogHealthcheckEvent:
			log.Printf("| UNIMPLEMENTED!\n")
		default:
			log.Printf("| unknown op code, echoing packet data\n")
			response, err := utils.MakePacket(r.Packet.OpCode, r.Packet.Data)
			if err != nil {
				log.Printf("| %s\n", err)
			}
			conn.WriteToUDP(response.Bytes(), r.RemoteAddr)
		}

	}
}

func Run() {
	setup.Keys()
	setup.Logger()

	address := setup.GetAddress()
	bufferLen, workerNum := setup.Config()

	buffer := make([]byte, bufferLen)
	requestsChannel := make(chan *Request)

	conn, err := net.ListenUDP("udp", address)
	if err != nil {
		log.Fatalf("| + %s\n", err)
	}
	defer conn.Close()

	for i := 0; i < workerNum; i++ {
		go handlePacket(conn, requestsChannel)
	}

	for {
		pLen, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("| error while reading packet : %s\n", err)
		}

		packet, err := utils.ParsePacket(buffer, pLen)
		if err != nil {
			log.Printf("| error while parsing packet : %s\n", err)
			continue
		}

		requestsChannel <- &Request{
			Packet:     packet,
			RemoteAddr: remoteAddr,
		}
	}
}
