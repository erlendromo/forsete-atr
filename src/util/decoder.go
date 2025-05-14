package util

import (
	"encoding/json"
	"io"
)

func DecodeJSON[T any](r io.Reader) (*T, error) {
	var target T
	if err := json.NewDecoder(r).Decode(&target); err != nil {
		return nil, err
	}

	return &target, nil
}
