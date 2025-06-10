# Simple BE Python & Golang

# Requirement

  - Golang 1.16.2 or higher
  - Python 3.7 or higher
  - Extract zip file and move to folder src project Go

### Note

If run auth-app golang error "gcc executable" you are running Ubuntu do:
```sh
$ apt-get install build-essential
```

If you want check deploy host you can change url API with 103.56.207.67

## Postman

If you want to check with Postman, you can import .json in the folder postman

## Case 1 : Auth App (Golang)

Go to the directory auth-app for setup and run app
```sh
$ cd auth-app
```

Setup dependencies
```sh
$ go get github.com/go-chi/chi/v5
$ go get github.com/stretchr/testify
$ go get github.com/newm4n/goornogo
$ go get github.com/mattn/go-sqlite3
$ go get github.com/dgrijalva/jwt-go
```

### Usage

Run application
```sh
$ go run main.py
```

### Makefile

build application
```sh
$ make build
```

run application
```sh
$ make run
```

unit test code
```sh
$ make test
```

clean go
```sh
$ make clean
```

docker build application
```sh
$ make docker-build
```

docker run application
```sh
$ make docker-run
```

### Check list example access API with curl

API register user
```sh
$ curl -d "phone=12345678&name=arfian&role=admin" -H "Content-Type: application/x-www-form-urlencoded" -X POST 103.56.207.67:2001/auth/register
```

API login
```sh
$ curl -d "phone=12345678&password=LiY4" -H "Content-Type: application/x-www-form-urlencoded" -X POST 103.56.207.67:2001/auth/login
```

API check jwt
```sh
$ curl -H "Content-Type: application/x-www-form-urlencoded" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7Im5hbWUiOiJhcmZpYW4iLCJwaG9uZSI6IjEyMzQ1Njc4Iiwicm9sZSI6ImFkbWluIiwidGltZXN0YW1wIjoiU3VuZGF5LCAwNC1BcHItMjEgMDI6MjM6NTAgVVRDIiwidXNlcm5hbWUiOiJhcmZpYW4ifSwiZXhwIjoxNjE3NzYyMzE4fQ.qV1tkmnj_-QK3ZrWWHdMbXihq-eTtIYPHiIGFtQvckA" -X GET 103.56.207.67:2001/auth/checkjwt
```

## Case 2 : Fetch App (Python)
Go to the directory fetch-app for setup and run app
```sh
$ cd fetch-app
```

### Installation

Install the dependencies. 
```sh
$ pip3 install -r requirements.txt
```

### Makefile  

unit test code  
```sh  
$ make test  
```  

run application  
```sh  
$ make run  
```  

docker build application
```sh
$ make docker-build
```

docker run application
```sh
$ make docker-run
```

### Check list example access API with curl

API get price USD
```sh
$ curl -H "Content-Type: application/x-www-form-urlencoded" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7Im5hbWUiOiJhcmZpYW4iLCJwaG9uZSI6IjEyMzQ1Njc4Iiwicm9sZSI6ImFkbWluIiwidGltZXN0YW1wIjoiU3VuZGF5LCAwNC1BcHItMjEgMDI6MjM6NTAgVVRDIiwidXNlcm5hbWUiOiJhcmZpYW4ifSwiZXhwIjoxNjE3NzYyMzE4fQ.qV1tkmnj_-QK3ZrWWHdMbXihq-eTtIYPHiIGFtQvckA" -X GET 103.56.207.67:2002/fetch/getprice
```

API Aggregate data
```sh
$ curl -H "Content-Type: application/x-www-form-urlencoded" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7Im5hbWUiOiJhcmZpYW4iLCJwaG9uZSI6IjEyMzQ1Njc4Iiwicm9sZSI6ImFkbWluIiwidGltZXN0YW1wIjoiU3VuZGF5LCAwNC1BcHItMjEgMDI6MjM6NTAgVVRDIiwidXNlcm5hbWUiOiJhcmZpYW4ifSwiZXhwIjoxNjE3NzYyMzE4fQ.qV1tkmnj_-QK3ZrWWHdMbXihq-eTtIYPHiIGFtQvckA" -X GET 103.56.207.67:2002/fetch/getaggregate
```

API check jwt
```sh
$ curl -H "Content-Type: application/x-www-form-urlencoded" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7Im5hbWUiOiJhcmZpYW4iLCJwaG9uZSI6IjEyMzQ1Njc4Iiwicm9sZSI6ImFkbWluIiwidGltZXN0YW1wIjoiU3VuZGF5LCAwNC1BcHItMjEgMDI6MjM6NTAgVVRDIiwidXNlcm5hbWUiOiJhcmZpYW4ifSwiZXhwIjoxNjE3NzYyMzE4fQ.qV1tkmnj_-QK3ZrWWHdMbXihq-eTtIYPHiIGFtQvckA" -X GET 103.56.207.67:2002/fetch/checkjwt
```
