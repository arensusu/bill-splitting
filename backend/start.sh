#!/bin/bash
set -e

source .env

/app/migrate -path /app/migration -database "postgres://$DATABASE_USER:$DATABASE_PASSWORD@$DATABASE_HOST:$DATABASE_PORT/$DATABASE_NAME?sslmode=disable" -verbose up

exec "$@"