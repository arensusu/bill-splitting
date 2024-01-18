#!/bin/bash
set -e

source .env

/app/migrate -path /app/migration -database "postgres://$DATABASE_USER:$DATABASE_PASSWORD@$DATABASE_HOST:5432/$DATABASE_NAME?sslmode=disable" -verbose up

exec "$@"