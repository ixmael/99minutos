#
#
#
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

RUN go build -o tracking_restapi ./cmd/restapi

#
#
#
FROM golang:1.25.7-alpine3.23

EXPOSE 8080

WORKDIR /app

RUN addgroup --system --gid 1000 golang \
    && adduser --system --uid 1000 golang

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY --chown=golang:golang .env.toml /app/.env.toml
COPY --chown=golang:golang internal/infrastructure/postgres/migrations /app/migrations
COPY --chown=golang:golang scripts/docker-entrypoint.sh /usr/bin/docker-entrypoint.sh
COPY --from=builder --chown=golang:golang /app/tracking_restapi /app/tracking_restapi

RUN chmod +x /app/tracking_restapi
RUN chmod +x /usr/bin/docker-entrypoint.sh

USER golang

ENTRYPOINT ["/usr/bin/docker-entrypoint.sh"]
CMD [ "/app/tracking_restapi" ]
# CMD [ "tail", "-f", "/dev/null" ]
