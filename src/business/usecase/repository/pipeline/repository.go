package pipeline

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/pipeline"
	querier "github.com/erlendromo/forsete-atr/src/business/usecase/querier/pipeline"
)

type PipelineRepository struct {
	querier querier.PipelineQuerier
}

func NewPipelineRepository(q querier.PipelineQuerier) *PipelineRepository {
	return &PipelineRepository{
		querier: q,
	}
}

func (p *PipelineRepository) RegisterPipeline(ctx context.Context, name, path string) (*pipeline.Pipeline, error) {
	return p.querier.RegisterPipeline(ctx, name, path)
}

func (p *PipelineRepository) PipelineByID(ctx context.Context, id int) (*pipeline.Pipeline, error) {
	return p.querier.PipelineByID(ctx, id)
}

func (p *PipelineRepository) PipelineByModel(ctx context.Context, textModelName string) (*pipeline.Pipeline, error) {
	return p.querier.PipelineByModel(ctx, textModelName)
}

func (p *PipelineRepository) PipelineByModels(ctx context.Context, lineModelName, textModelName string) (*pipeline.Pipeline, error) {
	return p.querier.PipelineByModels(ctx, lineModelName, textModelName)
}

func (p *PipelineRepository) ListPipelines(ctx context.Context) ([]*pipeline.Pipeline, error) {
	return p.querier.ListPipelines(ctx)
}

func (p *PipelineRepository) RegisterPipelineModel(ctx context.Context, pipelineID, modelID int) error {
	return p.querier.RegisterPipelineModel(ctx, pipelineID, modelID)
}
