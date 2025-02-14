package main

import (
	"github.com/erlendromo/forsete-atr/src/cmd/rest"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}

func main() {
	rest.StartRestService()
}
