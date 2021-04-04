FROM golang:1.16.2-alpine3.13

RUN mkdir /app

ADD . /app

RUN apk add --no-cache git
RUN apk add --no-cache sqlite-libs sqlite-dev
RUN apk add --no-cache build-base

WORKDIR /app

RUN go mod download

RUN go build -o main main.go

EXPOSE 2001

CMD ["/app/main"]