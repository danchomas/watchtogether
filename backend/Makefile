.PHONY: run test vet env-up env-down logs

run:
	go run ./cmd/server

test:
	go test -race ./cmd/... ./internal/...

vet:
	go vet ./cmd/... ./internal/...

env-up:
	docker compose -f ../docker-compose.yaml up --build -d

env-down:
	docker compose -f ../docker-compose.yaml down

logs:
	docker compose -f ../docker-compose.yaml logs -f
