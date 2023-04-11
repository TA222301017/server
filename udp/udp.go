package udp

import (
	"fmt"
	"log"
	"net"
	"time"

	"server/api/services"
	"server/models"
	gsetup "server/setup"
	"server/udp/setup"
	"server/udp/template"
	"server/udp/usecases"
	"server/udp/utils"
)

type Request struct {
	Packet     *template.BasePacket
	RemoteAddr *net.UDPAddr
}

func incomingPacketHandler(conn *net.UDPConn, request <-chan *Request) {
	var (
		res *template.BasePacket
		err error
	)

	for r := range request {
		log.Printf("| %s\n", r.RemoteAddr.String())
		switch r.Packet.OpCode {
		case template.KeyExchange:
			log.Printf("| 0x%x | KEY EXCHANGE\n", r.Packet.OpCode)
			res, err = usecases.KeyExchange(*r.Packet, r.RemoteAddr)

		case template.LogAccessEvent:
			log.Printf("| 0x%x | LOG ACCESS EVENT\n", r.Packet.OpCode)
			res, err = usecases.LogAccessEvent(*r.Packet)

		case template.LogRSSIEvent:
			log.Printf("| 0x%x | LOG RSSI EVENT\n", r.Packet.OpCode)
			res, err = usecases.LogRSSIEvent(*r.Packet)

		case template.LogHealthcheckEvent:
			log.Printf("| UNIMPLEMENTED!\n")
			res, err = utils.MakePacket(r.Packet.OpCode, r.Packet.Data, setup.PrivateKey)

		default:
			log.Printf("| unknown op code, echoing packet data\n")
			fmt.Println(r.Packet.OpCode)
			res, err = utils.MakePacket(r.Packet.OpCode, r.Packet.Data, setup.PrivateKey)
		}

		if res != nil {
			conn.WriteToUDP(res.Bytes(), r.RemoteAddr)
		}

		if err != nil {
			log.Printf("| %s\n", err)
		}
	}
}

func scanOutdatedAccessRules() {
	db := gsetup.DB

	var accessRules []models.AccessRule

	for {
		db.Find(&accessRules).Where("ends_at >= ?", time.Now())

		for _, accessRule := range accessRules {
			services.DeleteAccessRule(accessRule.ID)
		}

		time.Sleep(5 * time.Minute)
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

	go scanOutdatedAccessRules()

	for i := 0; i < workerNum; i++ {
		go incomingPacketHandler(conn, requestsChannel)
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
