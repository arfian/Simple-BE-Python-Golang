# efishery-task

# Requirement

  - Golang 1.16.2 or higher
  - Python 3.7 or higher
  - Extract zip file and move to folder src project Go

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

## Case 2 : Fetch App (Python)
Go to the directory fetch-app for setup and run app
```sh
$ cd fetch-app
```

### Installation

Install the dependencies. 
```sh
$ pip install -r requirements.txt
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