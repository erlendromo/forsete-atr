package pipeline

import (
	"os"
	"testing"
)

func setup() {
	_ = os.RemoveAll("assets")
	_ = os.MkdirAll("assets/pipelines", os.ModePerm)
}

func teardown() {
	_ = os.RemoveAll("assets")
}

type newPipelineTestCase struct {
	testName              string
	device                string
	filename              string
	segmentationModels    []string
	textRecognitionModels []string
	expectedPass          bool
}

var newPipelineTestCases = []newPipelineTestCase{
	{testName: "test1", device: "cpu", filename: "yolov9-lines-within-regions-1_TrOCR-norhand-v3", segmentationModels: []string{"yolov9-regions-1", "yolov9-lines-within-regions-1"}, textRecognitionModels: []string{"TrOCR-norhand-v3"}, expectedPass: true},
	{testName: "test2", device: "cpu", filename: "yolov9-lines-within-regions-1_TrOCR-norhand-v3", segmentationModels: []string{"yolov9-lines-within-regions-1"}, textRecognitionModels: []string{"TrOCR-norhand-v3"}, expectedPass: true},
	{testName: "test3", device: "cuda", filename: "yolov9-lines-within-regions-1_trocr-base-handwritten-hist-swe-2", segmentationModels: []string{"yolov9-lines-within-regions-1"}, textRecognitionModels: []string{"trocr-base-handwritten-hist-swe-2"}, expectedPass: true},
	{testName: "test4", device: "cuda", filename: "TrOCR-norhand-v3", segmentationModels: []string{""}, textRecognitionModels: []string{"TrOCR-norhand-v3"}, expectedPass: true},
	{testName: "test5", device: "cuda", filename: "yolov9-lines-within-regions-1", segmentationModels: []string{"yolov9-lines-within-regions-1"}, textRecognitionModels: []string{""}, expectedPass: true},

	{testName: "test6", device: "", filename: "yolov9-lines-within-regions-1_TrOCR-norhand-v3", segmentationModels: []string{"yolov9-lines-within-regions-1"}, textRecognitionModels: []string{"TrOCR-norhand-v3"}, expectedPass: false},
	{testName: "test7", device: "cpu", filename: "", segmentationModels: []string{"yolov9-lines-within-regions-1"}, textRecognitionModels: []string{"TrOCR-norhand-v3"}, expectedPass: false},
}

func TestNewPipeline(t *testing.T) {
	t.Run("Test new pipeline", testNewPipeline)
}
func testNewPipeline(t *testing.T) {
	for _, testCase := range newPipelineTestCases {
		if _, err := NewPipeline(testCase.device, testCase.filename); (err == nil) != testCase.expectedPass {
			t.Errorf("Test failed for file %s: expected error: %v, got error: %v", testCase.testName, testCase.expectedPass, err)
		}
	}
}

func TestCreateLocalYaml(t *testing.T) {
	setup()
	defer teardown()

	t.Run("Test create local yaml", testCreateLocalYaml)
}
func testCreateLocalYaml(t *testing.T) {
	for _, testCase := range newPipelineTestCases {
		pipeline, err := NewPipeline(testCase.device, testCase.filename)
		if err != nil {
			continue
		}

		for _, pathToModel := range testCase.segmentationModels {
			pipeline.AppendYoloStep(pathToModel)
		}

		for _, pathToModel := range testCase.textRecognitionModels {
			pipeline.AppendYoloStep(pathToModel)
		}

		pipeline.AppendOrderStep("OrderLines").AppendExportStep("json")

		if _, err := pipeline.CreateLocalYaml(); (err == nil) != testCase.expectedPass {
			t.Errorf("Test failed for file %s: expected error: %v, got error: %v", testCase.testName, testCase.expectedPass, err)
		}
	}
}
