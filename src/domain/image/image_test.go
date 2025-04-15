package image

import (
	"bytes"
	"fmt"
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

func setup() {
	_ = os.RemoveAll("assets")
	_ = os.MkdirAll("assets/images", os.ModePerm)
}

func teardown() {
	_ = os.RemoveAll("assets")
}

type NewImageTestCase struct {
	name         string
	file         multipart.File
	expectedPass bool
}

var newImageTestCases = []NewImageTestCase{
	{name: "xxx.png", file: fakeMultipartFile(), expectedPass: true},
	{name: "yyy.jpg", file: fakeMultipartFile(), expectedPass: true},
	{name: "zzz.jpeg", file: fakeMultipartFile(), expectedPass: true},

	{name: "aaa", file: fakeMultipartFile(), expectedPass: false},
	{name: "bbbgif", file: fakeMultipartFile(), expectedPass: false},
	{name: "ccc.pdf", file: fakeMultipartFile(), expectedPass: false},
	{name: "", file: fakeMultipartFile(), expectedPass: false},
	{name: "ddd.png", file: nil, expectedPass: false},
	{name: "", file: nil, expectedPass: false},
}

func TestNewImage(t *testing.T) {
	t.Run("New image test", testNewImage)
}
func testNewImage(t *testing.T) {
	for _, testCase := range newImageTestCases {
		_, err := NewImage(testCase.name, testCase.file)
		if (err == nil) != testCase.expectedPass {
			t.Errorf("Test failed for file %s: expected error: %v, got error: %v", testCase.name, testCase.expectedPass, err)
		}
	}
}

func TestCreateLocalImage(t *testing.T) {
	setup()
	defer teardown()

	t.Run("Create local image test", testCreateLocalImage)
}
func testCreateLocalImage(t *testing.T) {
	for _, testCase := range newImageTestCases {
		image, err := NewImage(testCase.name, testCase.file)
		if err != nil {
			continue
		}

		expectedImageName := fmt.Sprintf("assets/images/%s", testCase.name)
		localImageName, err := image.CreateLocalImage()
		if (err == nil) != testCase.expectedPass {
			t.Errorf("Test failed for file %s: expected error: %v, got error: %v", testCase.name, testCase.expectedPass, err)
		} else if expectedImageName != localImageName {
			t.Errorf("Test failed for file %s: expected name: %s, got name: %s", testCase.name, expectedImageName, localImageName)
		}
	}
}
