package utils

import (
	"fmt"
	"net"
	"os"
	"server/udp/template"
	"time"
)

var RemotePort string = "8888"

func SendUDPPacket(p *template.BasePacket, ipAddress string, port ...string) (*template.BasePacket, error) {
	lockAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", ipAddress, RemotePort))
	if len(port) == 1 {
		lockAddr, err = net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", ipAddress, port[0]))
	}

	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, lockAddr)
	if err != nil {
		println("Listen failed:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	if _, err := conn.Write(p.Bytes()); err != nil {
		return nil, err
	}

	responseBuffer := make([]byte, 2048)
	conn.SetReadDeadline(time.Now().Add(time.Second * 5))
	pLen, err := conn.Read(responseBuffer)
	if err != nil {
		return nil, err
	}

	return ParsePacket(responseBuffer, pLen)
}
