package modelstore

import (
	"bytes"
	"mime/multipart"
	"os"
	"testing"
)

func fakeMultipartFile() multipart.File {
	return &fakeFile{Reader: bytes.NewReader([]byte("fake image data"))}
}

type fakeFile struct {
	*bytes.Reader
}

func (f *fakeFile) Close() error {
	return nil
}

func setup(t *testing.T) {
	os.RemoveAll("assets")

	os.MkdirAll("assets/models/regionsegmentation/yolov9-regions-1", os.ModePerm)
	os.Create("assets/models/regionsegmentation/yolov9-regions-1/model.pt")

	os.MkdirAll("assets/models/linesegmentation/yolov9-lines-within-regions-1", os.ModePerm)
	os.Create("assets/models/linesegmentation/yolov9-lines-within-regions-1/model.pt")

	os.MkdirAll("assets/models/textrecognition/TrOCR-norhand-v3", os.ModePerm)
	os.Create("assets/models/textrecognition/TrOCR-norhand-v3/config.json")
	os.Create("assets/models/textrecognition/TrOCR-norhand-v3/generation_config.json")
	os.Create("assets/models/textrecognition/TrOCR-norhand-v3/merges.txt")
	os.Create("assets/models/textrecognition/TrOCR-norhand-v3/model.safetensors")
	os.Create("assets/models/textrecognition/TrOCR-norhand-v3/preprocessor_config.json")
	os.Create("assets/models/textrecognition/TrOCR-norhand-v3/special_tokens_map.json")
	os.Create("assets/models/textrecognition/TrOCR-norhand-v3/tokenizer.json")
	os.Create("assets/models/textrecognition/TrOCR-norhand-v3/tokenizer_config.json")
	os.Create("assets/models/textrecognition/TrOCR-norhand-v3/vocab.json")

	os.MkdirAll("assets/models/textrecognition/trocr-base-handwritten-hist-swe-2", os.ModePerm)
	os.Create("assets/models/textrecognition/trocr-base-handwritten-hist-swe-2/config.json")
	os.Create("assets/models/textrecognition/trocr-base-handwritten-hist-swe-2/generation_config.json")
	os.Create("assets/models/textrecognition/trocr-base-handwritten-hist-swe-2/merges.txt")
	os.Create("assets/models/textrecognition/trocr-base-handwritten-hist-swe-2/model.safetensors")
	os.Create("assets/models/textrecognition/trocr-base-handwritten-hist-swe-2/preprocessor_config.json")
	os.Create("assets/models/textrecognition/trocr-base-handwritten-hist-swe-2/special_tokens_map.json")
	os.Create("assets/models/textrecognition/trocr-base-handwritten-hist-swe-2/tokenizer.json")
	os.Create("assets/models/textrecognition/trocr-base-handwritten-hist-swe-2/tokenizer_config.json")
	os.Create("assets/models/textrecognition/trocr-base-handwritten-hist-swe-2/vocab.json")

	if err := GetModelstore().Initialize(); err != nil {
		t.Errorf("Test failed when initializing models, got error: %s", err.Error())
	}
}

func teardown() {
	_ = os.RemoveAll("assets")
}

type pathToModelTestCase struct {
	modelstore   *ModelStore
	modelName    string
	expectedPath string
}

func TestPathToModel(t *testing.T) {
	setup(t)
	defer teardown()

	t.Run("Test path to model", testPathToModel)
}
func testPathToModel(t *testing.T) {
	pathToModelTestCases := []pathToModelTestCase{
		{modelstore: GetModelstore(), modelName: "yolov9-lines-within-regions-1", expectedPath: "assets/models/linesegmentation/yolov9-lines-within-regions-1/model.pt"},
		{modelstore: GetModelstore(), modelName: "TrOCR-norhand-v3", expectedPath: "assets/models/textrecognition/TrOCR-norhand-v3"},

		{modelstore: GetModelstore(), modelName: "somethinginvalid", expectedPath: ""},
		{modelstore: GetModelstore(), modelName: "", expectedPath: ""},
	}

	for _, testCase := range pathToModelTestCases {
		path, found := testCase.modelstore.PathToModel(testCase.modelName)
		if (path == testCase.expectedPath) != found && path != "" {
			t.Errorf("Test failed for file %s: expected path: %s, got path: %s", testCase.modelName, testCase.expectedPath, path)
		}
	}
}

type modelsByTypeTestCase struct {
	modelstore         *ModelStore
	modelType          string
	expectedModelNames []string
	expectedPass       bool
}

func TestModelsByType(t *testing.T) {
	setup(t)
	defer teardown()

	t.Run("Test models by type", testModelsByType)
}
func testModelsByType(t *testing.T) {
	modelsByTypeTestCases := []modelsByTypeTestCase{
		{modelstore: GetModelstore(), modelType: "regionsegmentation", expectedModelNames: []string{"yolov9-regions-1"}, expectedPass: true},
		{modelstore: GetModelstore(), modelType: "linesegmentation", expectedModelNames: []string{"yolov9-lines-within-regions-1"}, expectedPass: true},
		{modelstore: GetModelstore(), modelType: "textrecognition", expectedModelNames: []string{"TrOCR-norhand-v3", "trocr-base-handwritten-hist-swe-2"}, expectedPass: true},
		{modelstore: GetModelstore(), modelType: "", expectedModelNames: []string{}, expectedPass: true},
	}

	for _, testCase := range modelsByTypeTestCases {
		models := testCase.modelstore.ModelsByType(testCase.modelType)

		if (len(models) == len(testCase.expectedModelNames)) != testCase.expectedPass {
			t.Errorf("Test failed: expected %d models, got %d models", len(testCase.expectedModelNames), len(models))
			continue
		}

		if testCase.expectedPass {
			actualNames := make([]string, len(models))
			for i, m := range models {
				actualNames[i] = m.Name
			}

			for _, expectedName := range testCase.expectedModelNames {
				found := false
				for _, actualName := range actualNames {
					if actualName == expectedName {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected model '%s' not found for type '%s'. Got: %v",
						expectedName, testCase.modelType, actualNames)
				}
			}
		}
	}
}

type addModelTestCase struct {
	modelstore   *ModelStore
	modelName    string
	modelType    string
	files        map[string]multipart.File
	expectedPass bool
}

func TestAddModel(t *testing.T) {
	setup(t)
	defer teardown()

	t.Run("Test add model", testAddModel)
}
func testAddModel(t *testing.T) {
	addModelTestCases := []addModelTestCase{
		{modelstore: GetModelstore(), modelName: "yolov9-some-test-model-1", modelType: "regionsegmentation", files: map[string]multipart.File{"model.pt": fakeMultipartFile()}, expectedPass: true},
		{modelstore: GetModelstore(), modelName: "yolov9-some-other-test-model-1", modelType: "linesegmentation", files: map[string]multipart.File{"model.pt": fakeMultipartFile()}, expectedPass: true},
		{modelstore: GetModelstore(), modelName: "TrOCR-some-test-model-1", modelType: "textrecognition", files: map[string]multipart.File{
			"config.json":              fakeMultipartFile(),
			"generation_config.json":   fakeMultipartFile(),
			"merges.txt":               fakeMultipartFile(),
			"model.safetensors":        fakeMultipartFile(),
			"preprocessor_config.json": fakeMultipartFile(),
			"special_tokens_map.json":  fakeMultipartFile(),
			"tokenizer.json":           fakeMultipartFile(),
			"tokenizer_config.json":    fakeMultipartFile(),
			"vocab.json":               fakeMultipartFile(),
		}, expectedPass: true},

		{modelstore: GetModelstore(), modelName: "yolov9-some-failing-test-model-1", modelType: "regionsegmentation", files: map[string]multipart.File{}, expectedPass: false},
		{modelstore: GetModelstore(), modelName: "yolov9-some-other-failing-test-model-1", modelType: "linesegmentation", files: map[string]multipart.File{}, expectedPass: false},
		{modelstore: GetModelstore(), modelName: "TrOCR-some-failing-test-model-1", modelType: "textrecognition", files: map[string]multipart.File{"model.safetensors": fakeMultipartFile()}, expectedPass: false},
	}

	for _, testCase := range addModelTestCases {
		err := testCase.modelstore.AddModel(testCase.modelName, testCase.modelType, testCase.files)
		if (err == nil) != testCase.expectedPass {
			t.Errorf("Test failed for model '%s', expected error: %v, got error: %v", testCase.modelName, testCase.expectedPass, err)
		}
	}
}
