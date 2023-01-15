package setup

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Logger() {
	// Naming log file
	logFileName := os.Getenv("UDP_LOG_FILE")
	if os.Getenv("UDP_LOG_FILE") != "" {
		if strings.HasSuffix(os.Getenv("UDP_LOG_FILE"), ".log") {
			logFileName = os.Getenv("UDP_LOG_FILE")
		} else {
			logFileName = os.Getenv("UDP_LOG_FILE") + ".log"
		}
	} else {
		logFileName = "udp.log"
	}

	// Opening log file
	logFile, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Failed to open log file : %v\n", err)
		os.Exit(1)
	}

	// Creating application logger
	logWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(logWriter)
	log.SetPrefix("[UDP] ")
	log.SetFlags(log.LstdFlags)
}
