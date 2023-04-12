package main

import (
	"server/models"
	"server/setup"
	"server/udp"
	"server/web"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	setup.Env()
	setup.Database()
	setup.ChannelServer()

	s := gocron.NewScheduler(time.UTC)
	s.Every(2).Days().Do(func() {
		db := setup.DB
		now := time.Now()

		db.Delete(&models.AccessLog{}, "timestamp <= ?", now)
		db.Delete(&models.RSSILog{}, "timestamp <= ?", now)
		db.Delete(&models.AccessRule{}, "ends_at <= ?", now)
	})

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
}
