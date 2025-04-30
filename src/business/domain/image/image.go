package image

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

var validFileTypes = []string{
	"png",
	"jpg",
	"jpeg",
}

type Image struct {
	ID         uuid.UUID `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	Format     string    `db:"format" json:"format"`
	Path       string    `db:"path" json:"-"`
	UploadedAt time.Time `db:"uploaded_at" json:"-"`
	DeletedAt  time.Time `db:"deleted_at" json:"-"`
	UserID     uuid.UUID `db:"user_id" json:"-"`
}

func (i *Image) CreateLocal(dst string, fileHeader *multipart.FileHeader) (string, error) {
	if err := i.checkFileHeader(fileHeader); err != nil {
		return "", err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}

	defer file.Close()

	if err := os.MkdirAll(dst, os.ModeDir); err != nil {
		return "", err
	}

	localFile, err := os.Create(fmt.Sprintf("%s/%s", strings.TrimRight(dst, "/"), fileHeader.Filename))
	if err != nil {
		return "", err
	}

	defer localFile.Close()

	if _, err := io.Copy(localFile, file); err != nil {
		return "", err
	}

	return localFile.Name(), nil
}

// If fileHeader is ok -> error = nil
func (i *Image) checkFileHeader(fileHeader *multipart.FileHeader) error {
	if len(fileHeader.Filename) < 5 {
		return fmt.Errorf("invalid filename '%s', should be atleast 5 characters", fileHeader.Filename)
	}

	sequences := strings.Split(fileHeader.Filename, ".")
	if len(sequences) != 2 {
		return fmt.Errorf("invalid filename '%s', should contain only one filetype", fileHeader.Filename)
	}

	format := strings.ToLower(strings.TrimPrefix(filepath.Ext(fileHeader.Filename), "."))
	if format == "" || !i.isValidFileType(format) {
		return fmt.Errorf("unsupported or missing file extension")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return err
	}

	defer file.Close()

	if file == nil {
		return fmt.Errorf("invalid file, cannot be nil")
	}

	return nil
}

func (i *Image) isValidFileType(format string) bool {
	for _, t := range validFileTypes {
		if t == format {
			return true
		}
	}
	return false
}
