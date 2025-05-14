package image

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"mime/multipart"
	"os"
	"time"

	"github.com/google/uuid"
)

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
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}

	defer file.Close()

	decodedImage, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	outImage, err := os.Create(fmt.Sprintf("%s/%s.png", i.Path, i.ID.String()))
	if err != nil {
		return err
	}

	defer outImage.Close()

	if err := png.Encode(outImage, decodedImage); err != nil {
		return err
	}

	return nil
}

func (i *Image) DeleteLocal() error {
	return os.Remove(fmt.Sprintf("%s/%s.%s", i.Path, i.ID.String(), i.Format))
}
