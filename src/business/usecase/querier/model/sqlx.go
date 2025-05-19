package model

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/model"
	"github.com/jmoiron/sqlx"
)

type SQLModelQuerier struct {
	db *sqlx.DB
}

func NewSQLModelQuerier(db *sqlx.DB) *SQLModelQuerier {
	return &SQLModelQuerier{
		db: db,
	}
}

func (q *SQLModelQuerier) RegisterModel(ctx context.Context, name, path string, modelTypeID int) (*model.Model, error) {
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

	var m model.Model
	err := q.db.QueryRowxContext(ctx, query, name, path, modelTypeID).StructScan(&m)

	return &m, err
}

func (q *SQLModelQuerier) ModelByID(ctx context.Context, id int) (*model.Model, error) {
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

	var m model.Model
	err := q.db.QueryRowxContext(ctx, query, id).StructScan(&m)

	return &m, err
}

func (q *SQLModelQuerier) ModelByName(ctx context.Context, name string) (*model.Model, error) {
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

	var m model.Model
	err := q.db.QueryRowxContext(ctx, query, name).StructScan(&m)

	return &m, err
}

func (q *SQLModelQuerier) ModelsByType(ctx context.Context, modelType string) ([]*model.Model, error) {
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

	rows, err := q.db.QueryxContext(ctx, query, modelType)
	if err != nil {
		return nil, err
	}

	var models []*model.Model
	for rows.Next() {
		var m model.Model
		if err := rows.StructScan(&m); err != nil {
			return nil, err
		}

		models = append(models, &m)
	}

	return models, err
}

func (q *SQLModelQuerier) AllModels(ctx context.Context) ([]*model.Model, error) {
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

	rows, err := q.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var models []*model.Model
	for rows.Next() {
		var m model.Model
		if err := rows.StructScan(&m); err != nil {
			return nil, err
		}

		models = append(models, &m)
	}

	return models, err
}
