run:
	@go run main.go

test:
	@go test -v ./... --race

composecpu:
	@docker compose -f docker-compose.yaml up --build -d

composecuda:
	@docker compose -f docker-compose.yaml -f docker/docker-compose-cuda.yaml up --build -d

composedown:
	@docker compose down --volumes --remove-orphans

attach:
	@docker compose logs -f forsete-atr

swag:
	@swag init -g src/api/router/router.go && swag fmt

.PHONY: run test composecpu composecuda composedown attach swag
