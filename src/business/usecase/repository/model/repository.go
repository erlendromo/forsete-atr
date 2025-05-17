package model

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/model"
	querier "github.com/erlendromo/forsete-atr/src/business/usecase/querier/model"
)

type ModelRepository struct {
	querier querier.ModelQuerier
}

func NewModelRepository(q querier.ModelQuerier) *ModelRepository {
	return &ModelRepository{
		querier: q,
	}
}

func (m *ModelRepository) RegisterModel(ctx context.Context, name, path string, modelTypeID int) (*model.Model, error) {
	return m.querier.RegisterModel(ctx, name, path, modelTypeID)
}

func (m *ModelRepository) ModelByID(ctx context.Context, id int) (*model.Model, error) {
	return m.querier.ModelByID(ctx, id)
}

func (m *ModelRepository) ModelByName(ctx context.Context, name string) (*model.Model, error) {
	return m.querier.ModelByName(ctx, name)
}

func (m *ModelRepository) ModelsByType(ctx context.Context, modelType string) ([]*model.Model, error) {
	return m.querier.ModelsByType(ctx, modelType)
}

func (m *ModelRepository) AllModels(ctx context.Context) ([]*model.Model, error) {
	return m.querier.AllModels(ctx)
}
