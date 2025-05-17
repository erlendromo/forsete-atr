package model

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/model"
)

type ModelQuerier interface {
	RegisterModel(ctx context.Context, name, path string, modelTypeID int) (*model.Model, error)
	ModelByID(ctx context.Context, id int) (*model.Model, error)
	ModelByName(ctx context.Context, name string) (*model.Model, error)
	ModelsByType(ctx context.Context, modelType string) ([]*model.Model, error)
	AllModels(ctx context.Context) ([]*model.Model, error)
}
