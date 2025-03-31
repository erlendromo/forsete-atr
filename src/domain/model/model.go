package model

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
	"sync"

	"github.com/erlendromo/forsete-atr/src/domain/model/trocrmodel"
	"github.com/erlendromo/forsete-atr/src/domain/model/yolomodel"
	"github.com/erlendromo/forsete-atr/src/util"
)

var requiredYoloFiles = []string{
	"model.pt",
}

var requiredTrOCRFiles = []string{
	"model.safetensors",
	"config.json",
	"generation_config.json",
	"merges.txt",
	"preprocessor_config.json",
	"special_tokens_map.json",
	"tokenizer.json",
	"tokenizer_config.json",
	"vocab.json",
}

var models map[string]Model

type Model interface {
	Name() string
	Path() string
	Type() string
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

func InitModels() error {
	if models != nil {
		return errors.New("models already initialized")
	}

	models = make(map[string]Model, 0)

	regionSegmentationEntries, err := os.ReadDir(fmt.Sprintf("%s/%s", util.MODELS, util.REGION_SEGMENTATION))
	if err != nil {
		return err
	}

	if len(regionSegmentationEntries) > 0 {
		if err := readEntries(regionSegmentationEntries, requiredYoloFiles, util.REGION_SEGMENTATION); err != nil {
			return err
		}
	}

	lineSegmentationEntries, err := os.ReadDir(fmt.Sprintf("%s/%s", util.MODELS, util.LINE_SEGMENTATION))
	if err != nil {
		return err
	}

	if len(lineSegmentationEntries) > 0 {
		if err := readEntries(lineSegmentationEntries, requiredYoloFiles, util.LINE_SEGMENTATION); err != nil {
			return err
		}
	}

	textRecognitionEntries, err := os.ReadDir(fmt.Sprintf("%s/%s", util.MODELS, util.TEXT_RECOGNITION))
	if err != nil {
		return err
	}

	if len(textRecognitionEntries) > 0 {
		if err := readEntries(textRecognitionEntries, requiredTrOCRFiles, util.TEXT_RECOGNITION); err != nil {
			return err
		}
	}

	return nil
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
		fmt.Printf("\n%sCREATE FILE ERROR%s\n%v",
			util.RED,
			util.RESET,
			err,
		)
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
				fmt.Printf("\n%sCREATE FILE ERROR%s\n%v",
					util.RED,
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

func readEntries(modelEntries []fs.DirEntry, requiredFiles []string, modelType string) error {
	var errList []error

	for _, modelDir := range modelEntries {
		if !modelDir.IsDir() {
			continue
		}

		modelName := modelDir.Name()
		modelPath := fmt.Sprintf("%s/%s/%s", util.MODELS, modelType, modelName)

		if err := checkRequiredFiles(modelPath, requiredFiles); err != nil {
			errList = append(errList, fmt.Errorf("%s: %v", modelPath, err))
			continue
		}

		switch modelType {
		case util.REGION_SEGMENTATION, util.LINE_SEGMENTATION:
			models[modelName] = yolomodel.NewYoloModel(
				modelName,
				fmt.Sprintf("%s/model.pt", modelPath),
				modelType,
			)
		case util.TEXT_RECOGNITION:
			models[modelName] = trocrmodel.NewTrOCRModel(
				modelName,
				modelPath,
				modelType,
			)
		default:
			errList = append(errList, fmt.Errorf("unknown model type: %s", modelType))
		}
	}

	if len(errList) > 0 {
		return fmt.Errorf("errors occurred: %v", errList)
	}

	return nil
}

func checkRequiredFiles(dir string, filenames []string) error {
	for _, filename := range filenames {
		filePath := fmt.Sprintf("%s/%s", dir, filename)

		if _, err := os.Stat(filePath); errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("missing file: %s", filename)
		}
	}

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
	if modelType != util.REGION_SEGMENTATION && modelType != util.LINE_SEGMENTATION && modelType != util.TEXT_RECOGNITION {
		return "", errors.New("invalid model type")
	}

	return fmt.Sprintf("%s/%s/%s", util.MODELS, modelType, modelName), nil
}
