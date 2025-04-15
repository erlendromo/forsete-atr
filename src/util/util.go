package util

import (
	"fmt"
	"os"
)

func MustGetEnv(key string) string {
	v, found := os.LookupEnv(key)
	if !found {
		panic(fmt.Errorf("Environment variable '%s' not set...", key))
	}

	return v
}
