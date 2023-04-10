APP_BIN = app/build/app

migrate:
	@./bin/migrate create -ext sql -dir migrations -seq -digits 8 $(NAME)

migrate.up:
	migrate -path migrations -database ${DATABASE_URL} up

.PHONY: swagger
swagger:
	swag init --parseDependency --parseInternal --parseDepth 1 -g ./cmd/main.go -o ./docs
