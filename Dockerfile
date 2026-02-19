FROM golang:1.25.7-alpine3.23 AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/

RUN go mod download

COPY cmd /app/cmd
COPY internal /app/internal

RUN go build -ldflags "-s -w" -o tracking_restapi cmd/restapi/main.go

FROM golang:1.25.7-alpine3.23

WORKDIR /app

COPY .env.toml /app/.env.toml
COPY --from=builder /app/tracking_restapi /app/tracking_restapi

ENTRYPOINT [ "/app/tracking_restapi" ]
