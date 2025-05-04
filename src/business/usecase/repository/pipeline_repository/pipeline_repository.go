package pipelinerepository

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/pipeline"
	"github.com/erlendromo/forsete-atr/src/database"
	"github.com/jmoiron/sqlx"
)

type PipelineRepository struct {
	db *sqlx.DB
}

func NewPipelineRepository(db *sqlx.DB) *PipelineRepository {
	return &PipelineRepository{
		db: db,
	}
}

func (p *PipelineRepository) PipelineByID(ctx context.Context, id int) (*pipeline.Pipeline, error) {
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

	return database.QueryRowx[pipeline.Pipeline](ctx, p.db, query, id)
}

func (p *PipelineRepository) AllPipelines(ctx context.Context) ([]*pipeline.Pipeline, error) {
	query := `
		SELECT
			id,
			name,
			path
		FROM
			"pipeline"
	`

	return database.Queryx[pipeline.Pipeline](ctx, p.db, query)
}

func (p *PipelineRepository) RegisterPipeline(ctx context.Context, name, path string) (*pipeline.Pipeline, error) {
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

	return database.QueryRowx[pipeline.Pipeline](ctx, p.db, query, name, path)
}

func (p *PipelineRepository) PipelineByModel(ctx context.Context, textModelName string) (*pipeline.Pipeline, error) {
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

	return database.QueryRowx[pipeline.Pipeline](ctx, p.db, query, textModelName)
}

func (p *PipelineRepository) PipelineByModels(ctx context.Context, lineModelName, textModelName string) (*pipeline.Pipeline, error) {
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

	return database.QueryRowx[pipeline.Pipeline](ctx, p.db, query, lineModelName, textModelName)
}

func (p *PipelineRepository) RegisterPipelineModel(ctx context.Context, pipelineID, modelID int) (int, error) {
	query := `
		INSERT INTO
			"pipeline_model" (pipeline_id, model_id)
		VALUES
			($1, $2)
	`

	return database.ExecuteContext(ctx, p.db, query, pipelineID, modelID)
}
