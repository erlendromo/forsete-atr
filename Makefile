run:
	@go run main.go

test:
	@go test -v ./... --race

composeup:
	@docker compose up --build -d

composedown:
	@docker compose down --volumes --remove-orphans

removetmp:
	@rm -rf tmp/images/* && rm -rf tmp/outputs/* && rm -rf tmp/yaml/*

.PHONY: run test composeup composedown removetmp
