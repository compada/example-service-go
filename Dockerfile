FROM golang:1.19 AS build

WORKDIR /usr/src/app

COPY go.* ./

RUN go mod download

COPY ./ ./

RUN go build -o /service ./cmd/graph/main.go

FROM gcr.io/distroless/base-debian10 AS deploy

WORKDIR /usr/src/app

COPY --from=build /service ./service

ENV PORT=8080
EXPOSE $PORT

USER nonroot:nonroot

ENTRYPOINT ["./service"]
