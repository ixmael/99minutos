#!/bin/sh

migrate \
    -path=/app/migrations \
    -database="$DATABASE_URL" \
    up

exec "$@"
