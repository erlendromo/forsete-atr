package output

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/output"
	"github.com/google/uuid"
)

type OutputQuerier interface {
	RegisterOutput(ctx context.Context, name, format, path string, imageID, userID uuid.UUID) (*output.Output, error)
	OutputByID(ctx context.Context, outputID, imageID, userID uuid.UUID) (*output.Output, error)
	OutputsByImageID(ctx context.Context, imageID, userID uuid.UUID) ([]*output.Output, error)
	UpdateOutputByID(ctx context.Context, confirmed bool, outputID, imageID, userID uuid.UUID) (*output.Output, error)
	DeleteOutputByID(ctx context.Context, outputID, imageID, userID uuid.UUID) error
	DeleteOutputsByImageID(ctx context.Context, imageID, userID uuid.UUID) error
	DeleteUserOutputs(ctx context.Context, userID uuid.UUID) error
}
