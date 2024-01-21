
run_compose:
	docker compose up -d

down_compose:
	docker compose down

run:
	@go run cmd/main.go

tidy:
	@go mod tidy