package modelstore

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
	"sync"

	"github.com/erlendromo/forsete-atr/src/domain/modelstore/model"
	"github.com/erlendromo/forsete-atr/src/util"
)

var modelstore *ModelStore

type ModelStore struct {
	models map[string]*model.Model
}

func GetModelstore() *ModelStore {
	if modelstore == nil {
		modelstore = &ModelStore{
			models: make(map[string]*model.Model),
		}
	}

	return modelstore
}

func (m *ModelStore) Initialize() error {
	if len(m.models) > 0 {
		return nil
	}

	for _, modelType := range []string{util.REGION_SEGMENTATION, util.LINE_SEGMENTATION, util.TEXT_RECOGNITION} {
		if err := m.readLocalEntries(modelType); err != nil {
			return err
		}
	}

	return nil
}

func (m *ModelStore) PathToModel(name string) (string, bool) {
	model, found := m.models[name]
	if found {
		return model.Path(), true
	}

	return "", false
}

func (m *ModelStore) ModelsByType(modelType string) []*model.Model {
	modelsResponse := make([]*model.Model, 0)

	for _, model := range m.models {
		if model.Type() == modelType {
			modelsResponse = append(modelsResponse, model)
		} else {
			continue
		}
	}

	return modelsResponse
}

func (m *ModelStore) AddModel(name, modelType string, files map[string]multipart.File) error {
	if err := m.checkModelType(modelType); err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s/%s", util.MODELS, modelType, name)
	if err := os.MkdirAll(path, os.ModeDir); err != nil {
		return err
	}

	for _, required := range m.requiredModelTypeFiles(modelType) {
		if _, ok := files[required]; !ok {
			return fmt.Errorf("missing required file: %s", required)
		}
	}

	var wg sync.WaitGroup

	for fileName, file := range files {
		wg.Add(1)
		go func(fileName string, file multipart.File) {
			defer file.Close()
			if err := m.createLocalFile(file, fmt.Sprintf("%s/%s", path, fileName)); err != nil {
				util.NewInternalErrorLog("CREATE FILE ERROR", err)
			}
			wg.Done()
		}(fileName, file)
	}

	wg.Wait()

	model, err := model.NewModel(name, modelType)
	if err != nil {
		return err
	}

	m.models[name] = model

	return nil
}

func (m *ModelStore) readLocalEntries(modelType string) error {
	if err := m.checkModelType(modelType); err != nil {
		return err
	}

	entriesDir := fmt.Sprintf("%s/%s", util.MODELS, modelType)
	entries, err := os.ReadDir(entriesDir)
	if err != nil {
		return err
	}

	if len(entries) <= 0 {
		return fmt.Errorf("no models present for type '%s'", modelType)
	}

	var errList []error

	for _, modelDir := range entries {
		if !modelDir.IsDir() {
			continue
		}

		name := modelDir.Name()
		path := fmt.Sprintf("%s/%s/%s", util.MODELS, modelType, name)

		requiredFiles := m.requiredModelTypeFiles(modelType)
		var requiredFileError error

		for _, filename := range requiredFiles {
			filePath := fmt.Sprintf("%s/%s", path, filename)

			if _, requiredFileError = os.Stat(filePath); errors.Is(requiredFileError, fs.ErrNotExist) {
				errList = append(errList, fmt.Errorf("%s: %v", path, requiredFileError))
				break
			}
		}

		if requiredFileError != nil {
			continue
		}

		model, err := model.NewModel(name, modelType)
		if err != nil {
			return err
		}

		m.models[name] = model
	}

	if len(errList) > 0 {
		return fmt.Errorf("errors occurred: %v", errList)
	}

	return nil
}

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

func (m *ModelStore) requiredModelTypeFiles(modelType string) []string {
	switch modelType {
	case util.REGION_SEGMENTATION, util.LINE_SEGMENTATION:
		return requiredYoloFiles
	case util.TEXT_RECOGNITION:
		return requiredTrOCRFiles
	default:
		return nil
	}
}

func (m *ModelStore) createLocalFile(file multipart.File, path string) error {
	defer file.Close()

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

func (m *ModelStore) checkModelType(modelType string) error {
	if modelType != util.REGION_SEGMENTATION &&
		modelType != util.LINE_SEGMENTATION &&
		modelType != util.TEXT_RECOGNITION {
		return fmt.Errorf("invalid modeltype '%s'", modelType)
	}

	return nil
}
