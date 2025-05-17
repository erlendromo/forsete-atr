package image

import (
	"context"
	"errors"
	"time"

	"github.com/erlendromo/forsete-atr/src/business/domain/image"
	"github.com/google/uuid"
)

type MockImageQuerier struct {
	images map[uuid.UUID]*image.Image
}

func NewMockImageQuerier() *MockImageQuerier {
	return &MockImageQuerier{
		images: make(map[uuid.UUID]*image.Image),
	}
}

func (q *MockImageQuerier) RegisterImage(ctx context.Context, name, format, path string, userID uuid.UUID) (*image.Image, error) {
	id := uuid.New()
	img := &image.Image{
		ID:         id,
		Name:       name,
		Format:     format,
		Path:       path,
		UploadedAt: time.Now().UTC(),
		UserID:     userID,
	}

	q.images[id] = img
	return img, nil
}

func (q *MockImageQuerier) ImageByID(ctx context.Context, imageID, userID uuid.UUID) (*image.Image, error) {
	img, ok := q.images[imageID]
	if !ok || img.UserID != userID || img.DeletedAt != nil {
		return nil, errors.New("image not found")
	}

	return img, nil
}

func (q *MockImageQuerier) ImagesByUserID(ctx context.Context, userID uuid.UUID) ([]*image.Image, error) {
	var result []*image.Image
	for _, img := range q.images {
		if img.UserID == userID && img.DeletedAt == nil {
			result = append(result, img)
		}
	}

	return result, nil
}

func (q *MockImageQuerier) DeleteImageByID(ctx context.Context, imageID, userID uuid.UUID) error {
	img, ok := q.images[imageID]
	if !ok || img.UserID != userID || img.DeletedAt != nil {
		return errors.New("image not found or already deleted")
	}

	now := time.Now().UTC()
	img.DeletedAt = &now

	return nil
}

func (q *MockImageQuerier) DeleteUserImages(ctx context.Context, userID uuid.UUID) error {
	now := time.Now().UTC()
	for _, img := range q.images {
		if img.UserID == userID && img.DeletedAt == nil {
			img.DeletedAt = &now
		}
	}

	return nil
}

func (q *MockImageQuerier) Seed(images []*image.Image) {
	for _, i := range images {
		q.images[i.ID] = i
	}
}
