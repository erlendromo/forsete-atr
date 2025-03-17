package model

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
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

func Models() []Model {
	localModels := make([]Model, 0)
	for _, model := range models {
		localModels = append(localModels, model)
	}

	return localModels
}

func ModelsByType(modelType string) []Model {
	modelsResponse := make([]Model, 0)

	for _, model := range models {
		if model.Type() == modelType {
			modelsResponse = append(modelsResponse, model)
		} else {
			continue
		}
	}

	return modelsResponse
}

func AddYoloModel(modelName, modelType string, modelFile multipart.File) error {
	defer modelFile.Close()

	var modelDirPath string
	switch modelType {
	case "regionsegmentation":
		modelDirPath = fmt.Sprintf("models/regionsegmentation/%s", modelName)
	case "linesegmentation":
		modelDirPath = fmt.Sprintf("models/linesegmentation/%s", modelName)
	default:
		return errors.New("invalid model type")
	}

	if err := os.MkdirAll(modelDirPath, os.ModeDir); err != nil {
		return err
	}

	modelPath := fmt.Sprintf("%s/model.pt", modelDirPath)

	localModelFile, err := os.Create(modelPath)
	if err != nil {
		return err
	}

	defer localModelFile.Close()

	if _, err := io.Copy(localModelFile, modelFile); err != nil {
		return err
	}

	models[modelName] = yolomodel.NewYoloModel(modelName, modelPath, modelType)

	return nil
}
