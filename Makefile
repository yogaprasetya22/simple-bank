include .envrc
MIGRATIONS_PATH = ./cmd/migrate/migrations

.PHONY: test
test:
	@go test -v ./...

.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	migrate -path db/migration -database "$(DB_ADDR)" -verbose up

.PHONY: migrate-down
migrate-down:
	migrate -path db/migration -database "$(DB_ADDR)" -verbose down

.PHONY: seed
seed: 
	@go run cmd/migrate/seed/main.go

.PHONY: gen-docs
gen-docs:
	@swag init -g ./api/main.go -d cmd,internal && swag fmt
sqlc:
	sqlc generate

.PHONY: postgres startdb createdb dropdb migrateup migratedown sqlc
