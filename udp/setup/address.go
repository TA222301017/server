package setup

import (
	"log"
	"net"
	"os"
	"strconv"
)

func GetAddress() *net.UDPAddr {
	ipStr := os.Getenv("APP_HOST")
	if ipStr == "" {
		log.Fatalf("Missing environment variable \"APP_HOST\"\n")
	}

	ip := net.ParseIP(ipStr)
	if ip == nil {
		log.Fatalf("Failed to parse IP address from \"APP_HOST\" : %s\n", ipStr)
	}

	portStr := os.Getenv("UDP_PORT")
	if portStr == "" {
		log.Fatalf("Missing environment variable \"UDP_PORT\"\n")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Failed to parse port from \"UDP_PORT\" : %s\n", portStr)
	}

	return &net.UDPAddr{
		IP:   ip,
		Port: port,
	}
}
