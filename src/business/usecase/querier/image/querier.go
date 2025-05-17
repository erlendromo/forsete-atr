package image

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/image"
	"github.com/google/uuid"
)

type ImageQuerier interface {
	RegisterImage(ctx context.Context, name, format, path string, userID uuid.UUID) (*image.Image, error)
	ImageByID(ctx context.Context, imageID, userID uuid.UUID) (*image.Image, error)
	ImagesByUserID(ctx context.Context, userID uuid.UUID) ([]*image.Image, error)
	DeleteImageByID(ctx context.Context, imageID, userID uuid.UUID) error
	DeleteUserImages(ctx context.Context, userID uuid.UUID) error
}
