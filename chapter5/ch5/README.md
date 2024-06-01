# Product API

API ini menyediakan operasi CRUD untuk mengelola data produk.

## Menjalankan Server

Untuk menjalankan server, gunakan perintah berikut di terminal:

```sh
go run challange5.go
```

## Endpoint
### Get All Data Products
- URL: `localhost:8080/products`
- Method: `GET`
- Description: Mengambil daftar semua produk.

Response:
```json
{
    "data": [
        {
            "id": 1,
            "name": "Product 1",
            "price": 100000
        },
        {
            "id": 2,
            "name": "Product 2",
            "price": 200000
        }
    ],
    "status": "success"
}
```


### Create Data Product

- URL: `localhost:8080/products`
- Method: `POST`
- Description: Menambahkan data produk baru.

Request body:
```json
{
    "name": "new name produk",
    "price": 200000
}

```


Response:
```json
{
    "data": {
        "id": 2,
        "name": "new name produk",
        "price": 200000
    },
    "status": "success"
}
```

### Update Data Product
- URL: `localhost:8080/product?id=1`
- Method: `PUT`
- Description: Mengupdate data produk.

Request body:
```json
{
    "name": "update name produk",
    "price": 300000
}

```


Response:
```json
{
    "data": {
        "id": 1,
        "name": "update name produk",
        "price": 300000
    },
    "status": "success"
}
```

### Delete Data Product
- URL: `localhost:8080/products?id=1`
- Method: `DELETE`
- Description: menghapus daftar produk berdasarkan idnya

Response:
```json
{
    "status": "success"
}
```
