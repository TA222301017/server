package main

import (
	"server/setup"
	"server/udp"
	"server/web"
	"sync"
)

// TODO : Pake transaction di usecase add/edit/delete access rule
// TODO : Tambahin script buat ngeset CRON job yg ngebackup db + ngeclear kepemilikan key kalo semua access rule udh kadaluarsa
// TODO : Tambahin kolom owner di tabel keys di admin panel servie
// TODO : Bikin endpoint get keys bisa difilter berdasarkan used ato nggak (WIP)
// TODO : Tambahin fitur search di select field personel/key/location di halaman2 admin panel (optional bgt)
// TODO : Tambahin script buat uninstall / update (optional bgt)

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
