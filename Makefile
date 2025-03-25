run:
	@go run main.go

test:
	@go test -v ./... --race

composecpu:
	@docker compose -f docker-compose.yaml up --build -d

composegpu:
	@docker compose -f docker-compose.yaml -f docker/docker-compose-gpu.yaml up --build -d

composedown:
	@docker compose down --volumes --remove-orphans

attach:
	@docker compose logs -f

removetmp:
	@rm -rf tmp/images/* && rm -rf tmp/outputs/* && rm -rf tmp/yaml/*

swag:
	@swag init -g src/api/router/router.go && swag fmt

.PHONY: run test composecpu composegpu composedown attach removetmp swag
