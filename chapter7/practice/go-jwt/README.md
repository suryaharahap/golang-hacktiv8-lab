# Challange 7

## Cara Setup Database Postgresql di Docker
Jalankan postgresql di docker

```bash
docker run --name chapter7 -e POSTGRES_USER=ch7user -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=ch7db -p 5432:5432 -d postgres:15
```


## Menjalankan Server

Untuk menjalankan server, gunakan perintah berikut di terminal:

```sh
cd chapter7/practice/go-jwt
go run main.go
```

PORT: 8083


REFERNSI blog:
https://www.suryahr.tech/blog/postgreql-with-docker
