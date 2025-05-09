package util

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func MustGetEnv(key string) string {
	v, found := os.LookupEnv(key)
	if !found {
		panic(fmt.Errorf("Environment variable '%s' not set...", key))
	}

	return v
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func CreateLocal(fileHeader *multipart.FileHeader, path, name, ext string) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}

	defer file.Close()

	localFile, err := os.Create(fmt.Sprintf("%s/%s.%s", strings.TrimRight(path, "/"), name, ext))
	if err != nil {
		return err
	}

	defer localFile.Close()

	if _, err := io.Copy(localFile, file); err != nil {
		return err
	}

	return nil
}

func NameAndExtFromFileHeader(fileHeader *multipart.FileHeader) (string, string) {
	ext := filepath.Ext(fileHeader.Filename)
	originalName := strings.TrimSuffix(fileHeader.Filename, ext)

	name := strings.ToLower(originalName)
	format := strings.ToLower(strings.TrimPrefix(ext, "."))

	return name, format
}
