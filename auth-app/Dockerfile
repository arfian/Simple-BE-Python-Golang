FROM golang:1.16.2-alpine3.13

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download

RUN go build -o main main.go

EXPOSE 2001

CMD ["/app/main"]