# User API Spec MyGram(app)

## Register User
- Type: `POST`
- Endpoint: `localhost:8080/api/v1/users`
- Request Body:
  ```json
  {
    "name": "surya",
    "email": "surya@gmail.com",
    "occupation": "Backend Engineer",
    "age": 25,
    "password": "password"
  }
  ```

- Response Body:
  ```json
  {
    "meta": {
        "message": "Account has been registered",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 10,
        "name": "surya",
        "occupation": "Backend Engineer",
        "email": "surya@gmail.com",
        "age": 25,
        "token": "tokentokentoken"
    }
  }
  ```
