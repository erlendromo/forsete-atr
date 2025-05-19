package pipeline

import (
	"context"
	"errors"

	"github.com/erlendromo/forsete-atr/src/business/domain/pipeline"
)

type MockPipelineQuerier struct {
	pipelines      map[int]*pipeline.Pipeline
	pipelineModels map[int][]int
	nextID         int
}

func NewMockPipelineQuerier(nextID int) *MockPipelineQuerier {
	return &MockPipelineQuerier{
		pipelines:      make(map[int]*pipeline.Pipeline),
		pipelineModels: make(map[int][]int),
		nextID:         nextID,
	}
}

func (m *MockPipelineQuerier) RegisterPipeline(ctx context.Context, name, path string) (*pipeline.Pipeline, error) {
	p := &pipeline.Pipeline{
		ID:   m.nextID,
		Name: name,
		Path: path,
	}

	m.pipelines[p.ID] = p
	m.nextID++

	return p, nil
}

func (m *MockPipelineQuerier) PipelineByID(ctx context.Context, id int) (*pipeline.Pipeline, error) {
	p, ok := m.pipelines[id]
	if !ok {
		return nil, errors.New("pipeline not found")
	}

	return p, nil
}

func (m *MockPipelineQuerier) PipelineByModel(ctx context.Context, textModelName string) (*pipeline.Pipeline, error) {
	// TODO implement
	return nil, errors.New("pipeline not found for model")
}

func (m *MockPipelineQuerier) PipelineByModels(ctx context.Context, lineModelName, textModelName string) (*pipeline.Pipeline, error) {
	// TODO implement
	return nil, errors.New("pipeline not found for models")
}

func (m *MockPipelineQuerier) ListPipelines(ctx context.Context) ([]*pipeline.Pipeline, error) {
	var result []*pipeline.Pipeline
	for _, p := range m.pipelines {
		result = append(result, p)
	}

	return result, nil
}

func (m *MockPipelineQuerier) RegisterPipelineModel(ctx context.Context, pipelineID, modelID int) error {
	if _, ok := m.pipelines[pipelineID]; !ok {
		return errors.New("pipeline not found")
	}

	if m.pipelineModels[pipelineID] == nil {
		m.pipelineModels[pipelineID] = make([]int, 0)
	}

	m.pipelineModels[pipelineID] = append(m.pipelineModels[pipelineID], modelID)
	return nil
}

func (m *MockPipelineQuerier) Seed(pipelines []*pipeline.Pipeline, pipelineModels []map[int][]int) {
	for _, p := range pipelines {
		m.pipelines[p.ID] = p
	}

	for _, pm := range pipelineModels {
		for pID, mIDs := range pm {
			m.pipelineModels[pID] = mIDs
		}
	}
}
