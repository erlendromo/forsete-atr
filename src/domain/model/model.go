package model

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"sync"

	"github.com/erlendromo/forsete-atr/src/domain/model/trocrmodel"
	"github.com/erlendromo/forsete-atr/src/domain/model/yolomodel"
	"github.com/erlendromo/forsete-atr/src/util"
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
				entries, err := os.ReadDir(modelDirName)
				if err != nil {
					return err
				}

				found := false
				for _, entry := range entries {
					if entry.Name() == "model.pt" {
						found = true
						break
					}
				}

				if !found {
					return errors.New("Yolo-model missing 'model.pt' file...")
				}

				models[modelDirName] = yolomodel.NewYoloModel(
					modelDirName,
					fmt.Sprintf("models/%s/%s/model.pt", entryDirName, modelDirName),
					entryDirName,
				)
			} else if entryDirName == "textrecognition" {
				// TODO add check for files here?

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

	modelDirPath, err := modelDirPath(modelName, modelType)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(modelDirPath, os.ModeDir); err != nil {
		return err
	}

	modelPath := fmt.Sprintf("%s/model.pt", modelDirPath)

	if err := createLocalFile(modelFile, modelPath); err != nil {
		return err
	}

	models[modelName] = yolomodel.NewYoloModel(modelName, modelPath, modelType)

	return nil
}

func AddTrOCRModel(modelName, modelType string, files map[string]multipart.File) error {
	modelDirPath, err := modelDirPath(modelName, modelType)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(modelDirPath, os.ModeDir); err != nil {
		return err
	}

	var wg *sync.WaitGroup

	for fileName, file := range files {
		wg.Add(1)
		go func(fileName string, file multipart.File) {
			defer file.Close()
			if err := createLocalFile(file, fmt.Sprintf("%s/%s", modelDirPath, fileName)); err != nil {
				fmt.Printf("\n%s%s%s\n%v",
					util.RED,
					"ERROR",
					util.RESET,
					err,
				)
			}
			wg.Done()
		}(fileName, file)
	}

	wg.Wait()

	models[modelName] = trocrmodel.NewTrOCRModel(modelName, modelDirPath, modelType)

	return nil
}

func createLocalFile(file multipart.File, path string) error {
	localFile, err := os.Create(path)
	if err != nil {
		return err
	}

	defer localFile.Close()

	if _, err := io.Copy(localFile, file); err != nil {
		return err
	}

	return nil
}

func modelDirPath(modelName, modelType string) (string, error) {
	switch modelType {
	case "regionsegmentation":
		return fmt.Sprintf("models/regionsegmentation/%s", modelName), nil
	case "linesegmentation":
		return fmt.Sprintf("models/linesegmentation/%s", modelName), nil
	case "textrecognition":
		return fmt.Sprintf("models/textrecognition/%s", modelName), nil
	default:
		return "", errors.New("invalid model type")
	}
}
