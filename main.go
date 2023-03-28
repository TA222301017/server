package main

import (
	"server/setup"
	"server/udp"
	"server/web"
	"sync"
)

// TODO : Tambahin script buat ngeset CRON job yg ngebackup db + ngeclear kepemilikan key kalo semua access rule udh kadaluarsa
// TODO : Tambahin fitur search di select field personel/key/location di halaman2 admin panel (optional bgt)
// TODO : Tambahin script buat uninstall / update (optional bgt)

func main() {
	setup.Env()
	setup.Database()
	setup.ChannelServer()

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
