# include .env
# export $(shell sed 's/=.*//' .env)

DB_ADDR := "postgres://admin:finuserdb@localhost:5433/finuserdb?sslmode=disable"
PROD_DB_ADDR := "postgres://admin:finuserdb@localhost:5433/

.PHONY: mup-user
mup:
	@echo "Running DB migration on: $(DB_ADDR)"
	migrate -database "$(DB_ADDR)" -path services/user-service/migrate/migrations up


.PHONY: mdown-user
mdown:
	@echo "Running DB migration on: $(DB_ADDR)"
	migrate -database "$(DB_ADDR)" -path services/user-service/migrate/migrations down

.PHONY: mupprod-user
mup:
	@echo "Running DB migration on: $(DB_ADDR)"
	migrate -database "$(PROD_DB_ADDR)" -path services/user-service/migrate/migrations up


.PHONY: mdownprod-user
mdown:
	@echo "Running DB migration on: $(DB_ADDR)"
	migrate -database "$(PROD_DB_ADDR)" -path services/user-service/migrate/migrations down

.PHONY: si
si:
	@echo "Generating Swagger docs..."
	swag init -g services/api-gateway/cmd/main.go --parseDependency --parseInternal --parseDepth 3
	@echo "Swagger docs generated successfully."

.PHONY: sf
sf:
	@echo "Formatting Swagger docs..."
	swag fmt

.PHONY: dkup
dkup:
	@echo "Starting Docker containers..."
	docker compose up

.PHONY: dkdown
dkdown:
	@echo "Stopping Docker containers..."
	docker compose down


PROTO_DIR := proto
PROTO_SRC := $(wildcard $(PROTO_DIR)/*.proto)
GO_OUT := .

.PHONY: generate-proto
generate-proto:
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(GO_OUT) \
		--go-grpc_out=$(GO_OUT) \
		$(PROTO_SRC)