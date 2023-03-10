package setup

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Logger() {
	path := filepath.Join(".", "logs")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

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
	logFileName = "./logs/" + logFileName

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
