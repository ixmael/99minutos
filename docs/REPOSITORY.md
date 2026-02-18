# Repository

This section provides information on how to interact with the repository.

## Generate migrations

```sh
# Generate a new migration
migrate create \
    -seq \
    -ext sql \
    --dir internal/infrastructure/postgres/migrations \
    migrationidentifier
```

## Execute migrations

```sh
# Apply the migrations
migrate \
    -database "postgres://postgres:postgres@127.0.0.1:5432/99minutos?sslmode=disable" \
    -path "internal/infrastructure/postgres/migrations" \
    up
```
