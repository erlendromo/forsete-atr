package image

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/image"
	querier "github.com/erlendromo/forsete-atr/src/business/usecase/querier/image"
	"github.com/google/uuid"
)

type ImageRepository struct {
	querier querier.ImageQuerier
}

func NewImageRepository(q querier.ImageQuerier) *ImageRepository {
	return &ImageRepository{
		querier: q,
	}
}

func (i *ImageRepository) RegisterImage(ctx context.Context, name, format, path string, userID uuid.UUID) (*image.Image, error) {
	return i.querier.RegisterImage(ctx, name, format, path, userID)
}

func (i *ImageRepository) ImageByID(ctx context.Context, imageID, userID uuid.UUID) (*image.Image, error) {
	return i.querier.ImageByID(ctx, imageID, userID)
}

func (i *ImageRepository) ImagesByUserID(ctx context.Context, userID uuid.UUID) ([]*image.Image, error) {
	return i.querier.ImagesByUserID(ctx, userID)
}

func (i *ImageRepository) DeleteImageByID(ctx context.Context, imageID, userID uuid.UUID) error {
	return i.querier.DeleteImageByID(ctx, imageID, userID)
}

func (i *ImageRepository) DeleteUserImages(ctx context.Context, userID uuid.UUID) error {
	return i.querier.DeleteUserImages(ctx, userID)
}
