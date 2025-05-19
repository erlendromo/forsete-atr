package model

import (
	"context"
	"errors"
	"strings"

	"github.com/erlendromo/forsete-atr/src/business/domain/model"
)

type MockModelQuerier struct {
	models     map[int]*model.Model
	modelTypes map[int]string
	nextID     int
}

func NewMockModelQuerier(nextID int) *MockModelQuerier {
	return &MockModelQuerier{
		models: make(map[int]*model.Model),
		modelTypes: map[int]string{
			1: "regionsegmentation",
			2: "linesegmentation",
			3: "textrecognition",
		},
		nextID: nextID,
	}
}

func (q *MockModelQuerier) RegisterModel(ctx context.Context, name, path string, modelTypeID int) (*model.Model, error) {
	modelType, ok := q.modelTypes[modelTypeID]
	if !ok {
		return nil, errors.New("invalid modelTypeID")
	}

	newModel := &model.Model{
		ID:          q.nextID,
		Name:        name,
		Path:        path,
		ModelTypeID: modelTypeID,
		ModelType:   modelType,
	}

	q.models[q.nextID] = newModel
	q.nextID++

	return newModel, nil
}

func (q *MockModelQuerier) ModelByID(ctx context.Context, id int) (*model.Model, error) {
	if model, ok := q.models[id]; ok {
		return model, nil
	}

	return nil, errors.New("not found")
}

func (q *MockModelQuerier) ModelByName(ctx context.Context, name string) (*model.Model, error) {
	for _, model := range q.models {
		if model.Name == name {
			return model, nil
		}
	}

	return nil, errors.New("not found")
}

func (q *MockModelQuerier) ModelsByType(ctx context.Context, modelType string) ([]*model.Model, error) {
	validModelType := false
	var result []*model.Model
	for _, model := range q.models {
		if strings.EqualFold(model.ModelType, modelType) {
			validModelType = true
			result = append(result, model)
		}
	}

	if !validModelType {
		return nil, errors.New("not found")
	}

	return result, nil
}

func (q *MockModelQuerier) AllModels(ctx context.Context) ([]*model.Model, error) {
	var result []*model.Model
	for _, model := range q.models {
		result = append(result, model)
	}

	return result, nil
}

func (q *MockModelQuerier) Seed(models []*model.Model) {
	for _, m := range models {
		q.models[m.ID] = m
	}
}
