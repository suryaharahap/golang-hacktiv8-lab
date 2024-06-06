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

- Response Body (Success): ✅
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

- Respon Body (Failed): ❌
  ```json
  {
    "meta": {
        "message": "Register account failed",
        "code": 422,
        "status": "error"
    },
    "data": {
        "errors": [
            "Key: 'RegisterUserInput.Name' Error:Field validation for 'Name' failed on the 'required' tag"
        ]
    }
  }
  ```


## Login User
- Type: `POST`
- Endpoint: `localhost:8080/api/v1/sessions`
- Request Body:
  ```json
  {
    "email": "suryahrp@gmail.com",
    "password": "password"
  }
  ```

- Response Body (Success): ✅
  ```json
  {
    "meta": {
        "message": "Successfuly loggedin",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 14,
        "name": "Surya Harahap",
        "occupation": "Backend Engineer",
        "email": "suryahrp@gmail.com",
        "age": 23,
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNH0.nzphv8gQ7XQMd-4XzalgMDVLORokX5VTR70boSosANE"
    }
  }
  ```

- Respon Body (Failed): ❌
  ```json
  {
    "meta": {
        "message": "Login failed",
        "code": 422,
        "status": "error"
    },
    "data": {
        "errors": "No user found on that email"
    }
  }
  ```

