package main

import (
	"os"
	"server/models"
	"server/setup"
	"server/udp"
	"server/web"
	"sync"
	"time"
)

func main() {
	setup.Env()
	setup.Database()
	setup.ChannelServer()
	args := os.Args

	if len(args) == 1 {
		wg := new(sync.WaitGroup)
		wg.Add(2)

		go func() {
			udp.Run()
			wg.Done()
		}()

		go func() {
			web.Run()
			wg.Done()
		}()

		wg.Wait()
	} else {
		cmd := args[1]
		if cmd == "clear_logs" {
			db := setup.DB
			now := time.Now()

			db.Delete(&models.AccessLog{}, "timestamp <= ?", now)
			db.Delete(&models.RSSILog{}, "timestamp <= ?", now)
		}
	}
}
