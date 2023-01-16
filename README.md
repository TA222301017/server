# Server
Implementasi dari subsistem Server yang mencangkup Admin Panel Service, Access Rule Service, dan Logging Service.

## Software yang Dibutuhkan
- Go (setidaknya versi 1.19.4)
- PostgreSQL (setidaknya versi 15.0)
- pgAdmin 4 (opsional, setidaknya versi 6)
- Make (opsional, setidaknya versi 3.81)

## Struktur Source Code
Struktur folder adalah sebagai berikut
```
server
├───api (source code terkait admin panel service)
│   ├───controllers
│   ├───middlewares
│   ├───services
│   ├───setup
│   ├───template
│   └───utils
├───build (hasil kompilasi program)
├───keys (tempat penyimpanan kunci-kunci kriptografik)
├───models (model tabel-tabel pada database)
├───setup (setup yang diperlukan sebelum program dijalankan)
├───udp (source code untuk access rule service dan logging service)
|   ├───setup
|   ├───template
|   ├───usecases
|   └───utils
├───main.go
├───Makefile
└───.env
```

## Konfigurasi Program
Konfigurasi program dapat diatur di file yang bernama `.env`. File tersebut harus dibuat sendiri. file tersebut harus mencantumkan variabel-variabel yang ada pada file `env.example`.

## Eksekusi Program
Program dapat dieksekusi dengan perintah berikut.
```shell
go run main.go
```

## Kompilasi Program
Program dapat dikompilasi dengan perintah berikut.
```shell
make all
```
Hasil kompilasi akan ada pada folder `build`. Folder tersebut dapat dihapus dengan perintah berikut.
```shell
make clean
```
