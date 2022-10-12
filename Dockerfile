FROM golang:1.19 AS build

WORKDIR /usr/src/app

COPY go.* ./

RUN go mod download

COPY ./ ./

RUN go build -o /service ./cmd/graph/main.go

FROM gcr.io/distroless/base-debian10 AS deploy

WORKDIR /usr/src/app

COPY --from=build /service ./service

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./service"]
