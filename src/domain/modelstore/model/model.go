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

	var path string
	switch modelType {
	case util.REGION_SEGMENTATION, util.LINE_SEGMENTATION:
		path = fmt.Sprintf("%s/%s/%s/%s/model.pt", util.ASSETS, util.MODELS, modelType, name)
	case util.TEXT_RECOGNITION:
		path = fmt.Sprintf("%s/%s/%s/%s", util.ASSETS, util.MODELS, modelType, name)
	default:
		return nil, fmt.Errorf("invalid modeltype '%s'", modelType)
	}

	return &Model{
		Name:      name,
		modelType: modelType,
		path:      path,
	}, nil
}
