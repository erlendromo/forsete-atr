package modelrepository

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/model"
	"github.com/erlendromo/forsete-atr/src/database"
	"github.com/jmoiron/sqlx"
)

type ModelRepository struct {
	db *sqlx.DB
}

func NewModelRepository(db *sqlx.DB) *ModelRepository {
	return &ModelRepository{
		db: db,
	}
}

func (m *ModelRepository) AllModels(ctx context.Context) ([]*model.Model, error) {
	query := `
		SELECT
			m.id,
			m.name,
			m.path,
			m.model_type_id,
			mt.type AS model_type
		FROM
			"model" m
		JOIN
			"model_type" mt ON m.model_type_id = mt.id
	`

	return database.Queryx[model.Model](ctx, m.db, query)
}

func (m *ModelRepository) ModelsByType(ctx context.Context, modelType string) ([]*model.Model, error) {
	query := `
        SELECT
            m.id,
            m.name,
            m.path,
            m.model_type_id,
            mt.type AS model_type
        FROM
            "model" m
        JOIN
            "model_type" mt ON m.model_type_id = mt.id
        WHERE
            mt.type = $1
    `

	models, err := database.Queryx[model.Model](ctx, m.db, query, modelType)
	if err != nil {
		return nil, err
	}

	if len(models) < 1 {
		return make([]*model.Model, 0), nil
	}

	return models, nil
}
