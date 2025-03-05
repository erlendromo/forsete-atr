package model

import (
	"errors"
	"fmt"
	"os"

	"github.com/erlendromo/forsete-atr/src/domain/model/trocrmodel"
	"github.com/erlendromo/forsete-atr/src/domain/model/yolomodel"
)

var models map[string]Model

type Model interface {
	Name() string
	Path() string
	Type() string
}

func InitModels() error {
	if models != nil {
		return errors.New("models already initialized")
	}

	models = make(map[string]Model)

	entriesDir, err := os.ReadDir("models")
	if err != nil {
		return err
	}

	for _, entryDir := range entriesDir {
		if !entryDir.IsDir() {
			continue
		}
		entryDirName := entryDir.Name()

		modelsDir, err := os.ReadDir(fmt.Sprintf("models/%s", entryDirName))
		if err != nil {
			return err
		}

		for _, modelDir := range modelsDir {
			if !modelDir.IsDir() {
				continue
			}

			modelDirName := modelDir.Name()

			if entryDirName == "linesegmentation" || entryDirName == "regionsegmentation" {
				models[modelDirName] = yolomodel.NewYoloModel(
					modelDirName,
					fmt.Sprintf("models/%s/%s/model.pt", entryDirName, modelDirName),
					entryDirName,
				)
			} else if entryDirName == "textrecognition" {
				models[modelDirName] = trocrmodel.NewTrOCRModel(
					modelDirName,
					fmt.Sprintf("models/%s/%s", entryDirName, modelDirName),
					entryDirName,
				)
			} else {
				return errors.New("invalid directory or file accessed")
			}
		}
	}

	return nil
}

func Path(name string) (string, bool) {
	model, found := models[name]
	if !found {
		return "", found
	}

	return model.Path(), found
}
