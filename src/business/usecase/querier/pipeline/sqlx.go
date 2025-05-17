package pipeline

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/pipeline"
	"github.com/jmoiron/sqlx"
)

type SQLPipelineQuerier struct {
	db *sqlx.DB
}

func NewSQLPipelineQuerier(db *sqlx.DB) *SQLPipelineQuerier {
	return &SQLPipelineQuerier{
		db: db,
	}
}

func (q *SQLPipelineQuerier) RegisterPipeline(ctx context.Context, name, path string) (*pipeline.Pipeline, error) {
	query := `
		INSERT INTO
			"pipeline" (name, path)
		VALUES
			($1, $2)
		RETURNING
			id,
			name,
			path
	`

	var p pipeline.Pipeline
	err := q.db.QueryRowxContext(ctx, query, name, path).StructScan(&p)
	return &p, err
}

func (q *SQLPipelineQuerier) PipelineByID(ctx context.Context, id int) (*pipeline.Pipeline, error) {
	query := `
		SELECT
			id,
			name,
			path
		FROM
			"pipeline"
		WHERE
			id = $1
	`

	var p pipeline.Pipeline
	err := q.db.QueryRowxContext(ctx, query, id).StructScan(&p)
	return &p, err
}

func (q *SQLPipelineQuerier) PipelineByModel(ctx context.Context, textModelName string) (*pipeline.Pipeline, error) {
	query := `
		SELECT
			p.id,
			p.name,
			p.path
		FROM
			pipeline p
		JOIN
			pipeline_model pm ON p.id = pm.pipeline_id
		JOIN
			model m ON m.id = pm.model_id AND m.name = $1
		LIMIT 1
	`

	var p pipeline.Pipeline
	err := q.db.QueryRowxContext(ctx, query, textModelName).StructScan(&p)
	return &p, err
}

func (q *SQLPipelineQuerier) PipelineByModels(ctx context.Context, lineModelName, textModelName string) (*pipeline.Pipeline, error) {
	query := `
		SELECT
			p.id,
			p.name,
			p.path
		FROM
			pipeline p
		JOIN
			pipeline_model pm1 ON p.id = pm1.pipeline_id
		JOIN
			model m1 ON m1.id = pm1.model_id AND m1.name = $1
		JOIN
			pipeline_model pm2 ON p.id = pm2.pipeline_id
		JOIN
			model m2 ON m2.id = pm2.model_id AND m2.name = $2
		LIMIT 1
	`

	var p pipeline.Pipeline
	err := q.db.QueryRowxContext(ctx, query, lineModelName, textModelName).StructScan(&p)
	return &p, err
}

func (q *SQLPipelineQuerier) ListPipelines(ctx context.Context) ([]*pipeline.Pipeline, error) {
	query := `
		SELECT
			id,
			name,
			path
		FROM
			"pipeline"
	`

	rows, err := q.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pipelines []*pipeline.Pipeline
	for rows.Next() {
		var p pipeline.Pipeline
		if err := rows.StructScan(&p); err != nil {
			return nil, err
		}
		pipelines = append(pipelines, &p)
	}

	return pipelines, nil
}

func (q *SQLPipelineQuerier) RegisterPipelineModel(ctx context.Context, pipelineID, modelID int) error {
	query := `
		INSERT INTO
			"pipeline_model" (pipeline_id, model_id)
		VALUES
			($1, $2)
	`

	_, err := q.db.ExecContext(ctx, query, pipelineID, modelID)
	return err
}
