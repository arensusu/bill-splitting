include .env

migrate-up:
	migrate -path db/migration -database "postgres://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):5432/$(DATABASE_NAME)?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migration -database "postgres://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):5432/$(DATABASE_NAME)?sslmode=disable" -verbose down

mockdb:
	mockgen -package mockdb -destination db/mock/store.go bill-splitting/db/sqlc Store