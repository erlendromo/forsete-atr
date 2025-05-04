package image

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/erlendromo/forsete-atr/src/util"
	"github.com/google/uuid"
)

var validFileTypes = []string{
	"png",
	"jpg",
	"jpeg",
}

// Image
//
//	@Summary		Image
//	@Description	Image containing id, name, format etc.
type Image struct {
	ID         uuid.UUID  `db:"id" json:"id"`
	Name       string     `db:"name" json:"name"`
	Format     string     `db:"format" json:"format"`
	Path       string     `db:"path" json:"-"`
	UploadedAt time.Time  `db:"uploaded_at" json:"-"`
	DeletedAt  *time.Time `db:"deleted_at" json:"-"`
	UserID     uuid.UUID  `db:"user_id" json:"-"`
}

func (i *Image) CreateLocal(fileHeader *multipart.FileHeader) error {
	if err := i.checkFileHeader(fileHeader); err != nil {
		return err
	}

	if err := util.CreateLocal(fileHeader, i.Path, i.ID.String(), i.Format); err != nil {
		return err
	}

	return nil
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
