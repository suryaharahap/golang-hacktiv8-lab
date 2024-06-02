# Challange 6


Pada challenge 6 ini membuat API CRUD data products dengan menggunakan framework `Gin` dan menggunakan database `Postgresql`, dibantu dengan ORM GORM.

## Requirements
Database akses menggunakan  `GORM` dan http frameworknya menggunakan `Gin`.

- Gin: `go get -u github.com/gin-gonic/gin`
- Gorm: `go get -u gorm.io/gorm`
- Postgresql: `go get -u gorm.io/driver/postgres`


## Cara Setup Database Postgresql di Docker
Jalankan postgresql di docker

```bash
docker run --name ch6-postgres -e POSTGRES_USER=chapter6 -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=chapter6-db -p 5432:5432 -d postgres:15
```

```bash
psql -h 127.0.0.1 -U chapter6 -d chapter6db
```

## Menjalankan Server

Untuk menjalankan server, gunakan perintah berikut di terminal:

```sh
cd chapter6/challange/
go run main.go
```


REFERNSI blog:
https://www.suryahr.tech/blog/postgreql-with-docker
