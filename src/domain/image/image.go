package image

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func ProcessImage(imageFile multipart.File, imageHeader *multipart.FileHeader) (string, error) {
	defer imageFile.Close()

	localImage, err := os.Create(fmt.Sprintf("tmp/images/%s", imageHeader.Filename))
	if err != nil {
		return "", err
	}

	defer localImage.Close()

	if _, err := io.Copy(localImage, imageFile); err != nil {
		return "", err
	}

	return localImage.Name(), nil
}
