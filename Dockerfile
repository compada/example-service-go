FROM golang:1.19

WORKDIR /usr/src/app

COPY go.* ./

RUN go mod download

COPY ./ ./
