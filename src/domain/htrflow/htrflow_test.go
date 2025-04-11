package htrflow

import (
	"os"
	"testing"
)

func setup() {
	os.RemoveAll("tmp")
	os.RemoveAll("scripts")

	os.MkdirAll("scripts", 0755)
	os.Create("scripts/htrflow.sh")

	os.MkdirAll("tmp/images", 0755)
	os.Create("tmp/images/testimage.png")

	os.MkdirAll("tmp/yaml", 0755)
	os.Create("tmp/yaml/testyaml.yaml")

	os.MkdirAll("tmp/outputs/images", 0755)
	os.Create("tmp/outputs/images/testimage.json")
}

func teardown() {
	os.RemoveAll("tmp")
	os.RemoveAll("scripts")
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
		{imagePath: "tmp/images/testimage.png", yamlPath: "tmp/yaml/testyaml.yaml", resultDst: "tmp/outputs/images/testimage.json", expectedPass: true},

		{imagePath: "tmp/images/invalid.png", yamlPath: "tmp/yaml/invalid.yaml", resultDst: "", expectedPass: false},
		{imagePath: "tmp/images/invalid.png", yamlPath: "", resultDst: "tmp/outputs/images/invalid.json", expectedPass: false},
		{imagePath: "tmp/images/invalid.png", yamlPath: "", resultDst: "", expectedPass: false},
		{imagePath: "", yamlPath: "tmp/yaml/invalid.yaml", resultDst: "tmp/outputs/images/invalid.json", expectedPass: false},
		{imagePath: "", yamlPath: "tmp/yaml/invalid.yaml", resultDst: "", expectedPass: false},
		{imagePath: "", yamlPath: "", resultDst: "tmp/outputs/images/invalid.json", expectedPass: false},
		{imagePath: "", yamlPath: "", resultDst: "", expectedPass: false},
	}

	for _, testCase := range runHTRflowTestCases {
		_, err := NewHTRflow(testCase.yamlPath, testCase.imagePath, testCase.resultDst).Run()
		if (err == nil) != testCase.expectedPass {
			t.Errorf("Test failed for htrflow, expected pass: %v, got error: %v", testCase.expectedPass, err)
		}
	}
}
