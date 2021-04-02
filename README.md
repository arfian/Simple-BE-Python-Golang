# efishery-task

# Requirement

  - Golang 1.16.2

### Installation Golang

Change Directory
```sh
$ cd auth-app
```

Setup dependencies
```sh
$ go get github.com/go-chi/chi/v5
$	go get github.com/stretchr/testify
$	go get github.com/newm4n/goornogo
$	go get github.com/mattn/go-sqlite3
```
	

Setup sqlite data structure
```sh
$ sqlite3 databases/efishery.db < setup.sql 
```

