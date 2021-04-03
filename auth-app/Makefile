.PHONY: test clean build

build:
	go mod download
	go build -o setup main.go

test:
	go get github.com/newm4n/goornogo
	go test ./... -v -covermode=count -coverprofile=coverage.out
	goornogo -c 20 -i coverage.out

clean: 
	go clean

run:
	go mod download
	go run main.go

docker-build: 
	docker build -t auth-app -f Dockerfile .

docker-run: 
	docker run -p 2001:2001 auth-app:latest