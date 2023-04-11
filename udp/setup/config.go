package setup

import (
	"os"
	"strconv"
)

func Config() (buffLen int, workerNum int) {
	var err error

	buffLen, err = strconv.Atoi(os.Getenv("UDP_BUFF_SIZE"))
	if err != nil {
		buffLen = 2048
	}

	workerNum, err = strconv.Atoi(os.Getenv("UDP_WORKER_NUM"))
	if err != nil {
		workerNum = 10
	}

	return buffLen, workerNum
}
