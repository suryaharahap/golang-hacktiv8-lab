# Final Project

## Cara Setup Database Postgresql di Docker
Jalankan postgresql di docker

```bash
docker run --name finalproject -e POSTGRES_USER=finalproject -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=finalproject_db -p 5432:5432 -d postgres:15
```


## Menjalankan Server

Untuk menjalankan server, gunakan perintah berikut di terminal:

```sh
cd final-project/mygram
go run main.go
```

PORT: 8083


REFERNSI blog:
https://www.suryahr.tech/blog/postgreql-with-docker

# ENDPOINT API - MyGram(app)
## Users
- url: `/users/register`
- Type: `POST`
- Request:
  body:
  ```json
  {
    "age": "integer",
    "email": "string",
    "password": "string",

  }
  ```
