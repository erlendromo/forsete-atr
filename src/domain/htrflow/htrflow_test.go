package htrflow

import (
	"os"
	"testing"
)

func setup() {
	os.RemoveAll("assets")
	os.RemoveAll("scripts")

	os.MkdirAll("assets/scripts", 0755)
	os.Create("assets/scripts/htrflow.sh")

	os.MkdirAll("assets/images", 0755)
	os.Create("assets/images/testimage.png")

	os.MkdirAll("assets/pipelines", 0755)
	os.Create("assets/pipelines/testyaml.yaml")

	os.MkdirAll("assets/outputs/images", 0755)
	os.Create("assets/outputs/images/testimage.json")
}

func teardown() {
	os.RemoveAll("assets")
	os.RemoveAll("assets/scripts")
}

type runHTRflowTestCase struct {
	imagePath    string
	yamlPath     string
	resultDst    string
	expectedPass bool
}

func TestRunHTRflow(t *testing.T) {
	setup()
	defer teardown()

	t.Run("Test run HTRflow", testRunHTRflow)
}
func testRunHTRflow(t *testing.T) {
	runHTRflowTestCases := []runHTRflowTestCase{
		{imagePath: "assets/images/testimage.png", yamlPath: "assets/pipelines/testyaml.yaml", resultDst: "assets/outputs/images/testimage.json", expectedPass: true},

		{imagePath: "assets/images/invalid.png", yamlPath: "assets/pipelines/invalid.yaml", resultDst: "", expectedPass: false},
		{imagePath: "assets/images/invalid.png", yamlPath: "", resultDst: "assets/outputs/images/invalid.json", expectedPass: false},
		{imagePath: "assets/images/invalid.png", yamlPath: "", resultDst: "", expectedPass: false},
		{imagePath: "", yamlPath: "assets/pipelines/invalid.yaml", resultDst: "assets/outputs/images/invalid.json", expectedPass: false},
		{imagePath: "", yamlPath: "assets/pipelines/invalid.yaml", resultDst: "", expectedPass: false},
		{imagePath: "", yamlPath: "", resultDst: "assets/outputs/images/invalid.json", expectedPass: false},
		{imagePath: "", yamlPath: "", resultDst: "", expectedPass: false},
	}

	for _, testCase := range runHTRflowTestCases {
		_, err := NewHTRflow(testCase.yamlPath, testCase.imagePath, testCase.resultDst).Run()
		if (err == nil) != testCase.expectedPass {
			t.Errorf("Test failed for htrflow, expected pass: %v, got error: %v", testCase.expectedPass, err)
		}
	}
}
