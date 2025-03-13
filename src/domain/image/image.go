package image

import (
	"io"
	"mime/multipart"
	"os"
)

func ProcessImage(imageFile multipart.File, imageHeader *multipart.FileHeader) (string, error) {
	localImage, err := os.Create("tmp/images/" + imageHeader.Filename)
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(localImage, imageFile); err != nil {
		return "", err
	}

	return localImage.Name(), nil
}
