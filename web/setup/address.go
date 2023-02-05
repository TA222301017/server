package setup

import "os"

func GetAddress() string {
	host := os.Getenv("APP_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	return host + ":" + port
}
