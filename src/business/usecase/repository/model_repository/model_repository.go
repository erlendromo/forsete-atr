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

	return database.Queryx[model.Model](ctx, m.db, query, modelType)
}

func (m *ModelRepository) ModelByID(ctx context.Context, id int) (*model.Model, error) {
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
        	m.id = $1
	`

	return database.QueryRowx[model.Model](ctx, m.db, query, id)
}

func (m *ModelRepository) ModelByName(ctx context.Context, name string) (*model.Model, error) {
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
        	m.name = $1
	`

	return database.QueryRowx[model.Model](ctx, m.db, query, name)
}

func (m *ModelRepository) RegisterModel(ctx context.Context, name, path string, model_type_id int) (*model.Model, error) {
	query := `
		WITH "inserted_model" AS
			(
				INSERT INTO
					"model" (name, path, model_type_id)
				VALUES
					($1, $2, $3)
				RETURNING
					id,
					name,
					path,
					model_type_id
			)
		SELECT
			im.id,
			im.name,
			im.path,
			im.model_type_id,
			mt.type AS model_type
		FROM
			"inserted_model" im
		JOIN
			"model_type" mt ON im.model_type_id = mt.id;
	`

	return database.QueryRowx[model.Model](ctx, m.db, query, name, path, model_type_id)
}
