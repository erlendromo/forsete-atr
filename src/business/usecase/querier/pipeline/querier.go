package pipeline

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/pipeline"
)

type PipelineQuerier interface {
	RegisterPipeline(ctx context.Context, name, path string) (*pipeline.Pipeline, error)
	PipelineByID(ctx context.Context, id int) (*pipeline.Pipeline, error)
	PipelineByModel(ctx context.Context, textModelName string) (*pipeline.Pipeline, error)
	PipelineByModels(ctx context.Context, lineModelName, textModelName string) (*pipeline.Pipeline, error)
	ListPipelines(ctx context.Context) ([]*pipeline.Pipeline, error)
	RegisterPipelineModel(ctx context.Context, pipelineID, modelID int) error
}
