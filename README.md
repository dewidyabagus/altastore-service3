# Alta Store Back-End Microservices Service 3

Alta Store Back-End (AS BE) adalah layanan web service (RESTful API) untuk kebutuhan pemrosesan data layanan Toko Online dimana fitur yang diberikan mulai dari registrasi, login pelanggan dan admin, master produk dan kategori produk, keranjang belanja, metode pembayaran yang menggunakan payment gateway Midtrans dan menggunakan sistem autentikasi JWT.

## Fitur Pada Service 3

- Keranjang belanjaan (Shopping Cart)
- Pembayaran (Checkout & Payment)

## Teknologi dan Arsitektur

`Teknologi` dan `Arsitektur` pengembangan sistem back-end mulai dari bahasa pemrograman, database server hingga infrastruktur yang digunakan.

- **_Development Tools_**

  - Golang (Program Language)
  - Echo Framework (Web Framework)
  - PostgreSQL (Database Server)
  - MongoDB (Database Logging)

- **_API Doc & Tester_**

  - Swagger
  - Postman

- **_Container_**

  - Docker

- **_Arsitektur_**
  - Heksagonal

## Memulai Web Service API

- Menjalankan **_Web Service API_** pada lingkungan pengembangan
  Sebelum menjalankan **_Web Service API_** pastikan anda sudah menduplikasi file `.env_example` dan menggantinya menjadi `.env`.
  File tersebut digunakan untuk menyimpan konfigurasi aplikasi, alamat database dan secret JWT.

  Untuk menjalankan **_Web Service API_** ketikan perintah ini di terminal anda:

  ```console
  go run .
  ```

  Anda dapat melihatnya di `http://localhost:8000` secara default sesuai dengan konfigurasi `.env`

- Menjalankan **_Web Service API_** dengan **_Docker_** <br>
  Kemudahan menjalankan **_Web Service API_** dengan **_Docker_**. Anda akan mendapatkan kemudahan dalam menyiapkan **_Web Service API_**
  dengan menggunakan docker tanpa harus melakukan instalasi `go language`, `postgresql` dan `mongodb`, cukup melakukan
  [install `docker`](https://docs.docker.com/engine/install/) dan jalankan perintah dibawah ini. Docker akan menyiapkan
  semua keperluan yang dibutuhkan.

  - Menggunakan Sistem Operasi Windows

    ```console
    windows-docker-compose.bat
    ```

  - Menggunakan Sistem Operasi Linux <br>
    Pastikan Anda sudah [menginstall docker compose](https://docs.docker.com/compose/install/)

    ```console
    /bin/bash linux-docker-compose.sh
    ```

  **_Docker_** akan melakukan build image web service sesuai dengan file konfigurasi `.env`,
  pastikan nama host `postgres` dan `mongodb` sesuai dengan nama container **_"secara default sudah sama"_**.<br>
  **_Docker-compose_** akan membuatkan container sesuai dengan setup lingkuan variabel dan pastikan kembali
  konfigurasi lingkungan variabel sudah sesuai **_"secara default sudah sama"_**

  Anda dapat melihatnya di `http://localhost:8000` secara default sesuai dengan konfigurasi `.env`

## Melakukan Request Web Service API

Kami menyiapkan yang terbaik untuk Anda. Untuk melihat dokumentasi **_Web Service API_** sudah kami siapkan di file `AltaStore.yaml`.
Kami juga menyiapkan collection **_Postman_** untuk pengujian dan permintaan ke **_Web Service API_** di file `BE Alta Store Service 3.json`
jika belum memiliki **_Postman_** Anda dapat [download dan install](https://www.postman.com/downloads/) Postman terlebih dahulu.

## Hasil Unit Testing

Kami memastikan apa yang kami bagikan sudah melalui proses dan memastikan kualitas yang terbaik untuk Anda, kami sudah melakukan
pengujian dengan cara melakukan **_Unit Testing_** dan hasilnya dapat dilihat di https://sonarcloud.io/project/overview?id=dewidyabagus_altastore-service3 .
Kami mengharapkan kritik dan saran untuk celah keamanan, bug dan perbaikan aplikasi untuk lebih baik, kami akan terus memperbarui status pengujian terakhir
pada link yang sudah dibagikan diatas.

## Repository Service Lainnya

- [Alta Store Back-End Microservices Service 1](https://github.com/Yap0894/altastore-service1)
- [Alta Store Back-End Microservices Service 2](https://github.com/dewidyabagus/altastore-service2)

## Kontak

Kami sangat terbuka untuk berdiskusi, menerima kritik dan saran. Anda dapat menghubungi kami:

- Widya Ade Bagus - https://www.linkedin.com/in/widya-ade-bagus-3a660716b/
- Alexander Yap - https://www.linkedin.com/in/alexander-yap-a14015169/
