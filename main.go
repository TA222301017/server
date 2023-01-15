package main

import (
	"server/api"
	"server/setup"
	"server/udp"
	"sync"
)

func main() {
	setup.Env()
	setup.Database()

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		udp.Run()
		wg.Done()
	}()

	go func() {
		api.Run()
		wg.Done()
	}()

	wg.Wait()
}
