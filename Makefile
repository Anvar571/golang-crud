- include .env
DB_URL=postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable

run_compose:
	sudo docker compose up -d

down_compose:
	sudo docker compose down

run:
	go run cmd/main.go

tidy:
	go mod tidy

migrateup:
	migrate -database "$(DB_URL)" -path database/migrations up

migratedown:
	migrate -database "$(DB_URL)" -path database/migrations down