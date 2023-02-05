package main

import (
	"server/setup"
	"server/udp"
	"server/web"
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
		web.Run()
		wg.Done()
	}()

	wg.Wait()
}
