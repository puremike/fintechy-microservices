include .env
export $(shell sed 's/=.*//' .env)

.PHONY: mup-user
mup:
	@echo "Running DB migration on: $(DB_ADDR)"
	migrate -database "postgres://admin:finuserdb@localhost:5433/finuserdb?sslmode=disable" -path services/user-service/migrate/migrations up


.PHONY: mdown-user
mdown:
	@echo "Running DB migration on: $(DB_ADDR)"
	migrate -database "postgres://admin:finuserdb@localhost:5433/finuserdb?sslmode=disable" -path services/user-service/migrate/migrations down

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