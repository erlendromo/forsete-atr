package output

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/output"
	querier "github.com/erlendromo/forsete-atr/src/business/usecase/querier/output"
	"github.com/google/uuid"
)

type OutputRepository struct {
	querier querier.OutputQuerier
}

func NewOutputRepository(q querier.OutputQuerier) *OutputRepository {
	return &OutputRepository{
		querier: q,
	}
}

func (o *OutputRepository) RegisterOutput(ctx context.Context, name, format, path string, imageID, userID uuid.UUID) (*output.Output, error) {
	return o.querier.RegisterOutput(ctx, name, format, path, imageID, userID)
}

func (o *OutputRepository) OutputByID(ctx context.Context, outputID, imageID, userID uuid.UUID) (*output.Output, error) {
	return o.querier.OutputByID(ctx, outputID, imageID, userID)
}

func (o *OutputRepository) OutputsByImageID(ctx context.Context, imageID, userID uuid.UUID) ([]*output.Output, error) {
	return o.querier.OutputsByImageID(ctx, imageID, userID)
}

func (o *OutputRepository) UpdateOutputByID(ctx context.Context, confirmed bool, outputID, imageID, userID uuid.UUID) (*output.Output, error) {
	return o.querier.UpdateOutputByID(ctx, confirmed, outputID, imageID, userID)
}

func (o *OutputRepository) DeleteOutputByID(ctx context.Context, outputID, imageID, userID uuid.UUID) error {
	return o.querier.DeleteOutputByID(ctx, outputID, imageID, userID)
}

func (o *OutputRepository) DeleteOutputsByImageID(ctx context.Context, imageID, userID uuid.UUID) error {
	return o.querier.DeleteOutputsByImageID(ctx, imageID, userID)
}

func (o *OutputRepository) DeleteUserOutputs(ctx context.Context, userID uuid.UUID) error {
	return o.querier.DeleteUserOutputs(ctx, userID)
}
