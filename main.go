package main

import (
	"github.com/erlendromo/forsete-atr/src/cmd"
	"github.com/erlendromo/forsete-atr/src/domain/modelstore"
)

func init() {
	if err := modelstore.GetModelstore().Initialize(); err != nil {
		panic(err)
	}
}

func main() {
	cmd.StartService()
}
