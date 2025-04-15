package image

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

type Image struct {
	name           string
	file           multipart.File
	localImagePath string
}

func NewImage(name string, file multipart.File) (*Image, error) {
	if file == nil {
		return nil, errors.New("file cannot be nil")
	}

	defer file.Close()

	sequences := strings.Split(name, ".")
	if len(sequences) != 2 || len(name) < 5 {
		return nil, fmt.Errorf(
			"invalid filename '%s', should be atleast 5 characters and contain one filetype",
			name,
		)
	}

	filetype := sequences[len(sequences)-1]
	if !strings.Contains(filetype, "png") &&
		!strings.Contains(filetype, "jpg") &&
		!strings.Contains(filetype, "jpeg") {
		return nil, fmt.Errorf("unsupported filetype '%s'", filetype)
	}

	return &Image{
		name:           name,
		file:           file,
		localImagePath: fmt.Sprintf("assets/images/%s", name),
	}, nil
}

func (i *Image) CreateLocalImage() (string, error) {
	localImageFile, err := os.Create(i.localImagePath)
	if err != nil {
		return "", err
	}

	defer i.file.Close()
	defer localImageFile.Close()

	if _, err := io.Copy(localImageFile, i.file); err != nil {
		return "", err
	}

	return localImageFile.Name(), nil
}
