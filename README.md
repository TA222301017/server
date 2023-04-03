# Server

Implementasi dari subsistem Server yang mencangkup Admin Panel Service, Access Rule Service, dan Logging Service. Dokumentasi API Admin Panel Service dapat dilihat di [link ini](https://documenter.getpostman.com/view/11921205/2s8YmNPhAE).

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
├───logs (tempat penyimpanan file-file log)
├───models (model tabel-tabel pada database)
├───release (hasil cross compilation)
├───setup (setup yang diperlukan sebelum program dijalankan)
├───udp (source code untuk access rule service dan logging service)
|   ├───setup
|   ├───template
|   ├───usecases
|   └───utils
├───web (source code untuk web server)
├───main.go
├───Makefile
└───.env
```

## Konfigurasi Program

Konfigurasi program dapat diatur di file yang bernama `.env`. File tersebut harus dibuat sendiri. file tersebut harus mencantumkan variabel-variabel yang ada pada file `env.example`.

## Eksekusi Program

Program dapat dieksekusi dengan perintah berikut.

```shell
$ go run main.go
```

## Kompilasi Program

Program dapat dikompilasi dengan perintah berikut.

```shell
$ make build
```

Hasil kompilasi akan ada pada folder `build`. Folder tersebut dapat dihapus dengan perintah berikut.

```shell
$ make clean
```

## Instalasi

Rilis terbaru program ini beserta rilis terbaru [Admin Panel](https://github.com/TA222301017/admin-panel-js) untuk program ini dapat di-install pada sebuah komputer linux dengan perintah berikut.

```shell
$ curl https://raw.githubusercontent.com/TA222301017/server/main/install.sh > install.sh && source ./install.sh
```

Pastikan Anda sudah membuat file `.env` seperti pada file `.env.example` sebelum melakukan instalasi.
