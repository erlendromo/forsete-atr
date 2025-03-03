run:
	@go run main.go

test:
	@go test -v ./... --race

composeup:
	@docker compose up --build -d

composedown:
	@docker compose down --volumes --remove-orphans

.PHONY: run test composeup composedown
