package model

import (
	"fmt"
	"strings"

	"github.com/erlendromo/forsete-atr/src/util"
)

type Model struct {
	Name      string `json:"name"`
	modelType string
	path      string
}

func (m *Model) Type() string {
	return m.modelType
}

func (m *Model) Path() string {
	return m.path
}

func NewModel(name, modelType string) (*Model, error) {
	if len(name) < 5 || strings.ContainsAny(name, "*)(/&%$#!:;><@`´+±€£™©∞§|[]≈…‚") {
		return nil, fmt.Errorf("invalid model name '%s', cannot contain special characters except dash and underscore", name)
	}

	path := fmt.Sprintf("%s/%s/%s", util.MODELS, modelType, name)
	if modelType == util.REGION_SEGMENTATION || modelType == util.LINE_SEGMENTATION {
		path = fmt.Sprintf("%s/model.pt", path)
	}

	return &Model{
		Name:      name,
		modelType: modelType,
		path:      path,
	}, nil
}
