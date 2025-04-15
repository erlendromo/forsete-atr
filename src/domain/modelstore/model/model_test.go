package model

import "testing"

type newModelTestCase struct {
	name         string
	modelType    string
	expectedPass bool
}

var newModelTestCases = []newModelTestCase{
	{name: "yolov9-regions-1", modelType: "regionsegmentation", expectedPass: true},
	{name: "yolov9-lines-within-regions-1", modelType: "linesegmentation", expectedPass: true},
	{name: "TrOCR-norhand-v3", modelType: "textrecognition", expectedPass: true},

	{name: "xxx", modelType: "", expectedPass: false},
	{name: "", modelType: "textrecognition", expectedPass: false},
	{name: "shouldFail", modelType: "someInvalidModelType", expectedPass: false},
	{name: "", modelType: "", expectedPass: false},
}

func TestNewModel(t *testing.T) {
	t.Run("Test new model", testNewModel)
}
func testNewModel(t *testing.T) {
	for _, testCase := range newModelTestCases {
		if _, err := NewModel(testCase.name, testCase.modelType); (err == nil) != testCase.expectedPass {
			t.Errorf("Test failed for file %s: expected error: %v, got error: %v", testCase.name, testCase.expectedPass, err)
		}
	}
}
