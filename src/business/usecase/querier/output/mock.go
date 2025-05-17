package output

import (
	"context"
	"errors"
	"time"

	"github.com/erlendromo/forsete-atr/src/business/domain/output"
	"github.com/google/uuid"
)

type MockOutputQuerier struct {
	outputs map[uuid.UUID]*output.Output
	images  map[uuid.UUID]uuid.UUID
}

func NewMockOutputQuerier() *MockOutputQuerier {
	return &MockOutputQuerier{
		outputs: make(map[uuid.UUID]*output.Output),
		images:  make(map[uuid.UUID]uuid.UUID),
	}
}

func (q *MockOutputQuerier) RegisterOutput(ctx context.Context, name, format, path string, imageID, userID uuid.UUID) (*output.Output, error) {
	if owner, ok := q.images[imageID]; !ok || owner != userID {
		return nil, errors.New("unauthorized or missing image")
	}

	id := uuid.New()
	o := &output.Output{
		ID:        id,
		Name:      name,
		Format:    format,
		Path:      path,
		ImageID:   imageID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Confirmed: false,
	}

	q.outputs[id] = o
	return o, nil
}

func (q *MockOutputQuerier) OutputByID(ctx context.Context, outputID, imageID, userID uuid.UUID) (*output.Output, error) {
	o, ok := q.outputs[outputID]
	if !ok || o.DeletedAt != nil || o.ImageID != imageID {
		return nil, errors.New("output not found")
	}

	if owner, ok := q.images[o.ImageID]; !ok || owner != userID {
		return nil, errors.New("unauthorized or missing image")
	}

	return o, nil
}

func (q *MockOutputQuerier) OutputsByImageID(ctx context.Context, imageID, userID uuid.UUID) ([]*output.Output, error) {
	if owner, ok := q.images[imageID]; !ok || owner != userID {
		return nil, errors.New("unauthorized or missing image")
	}

	var results []*output.Output
	for _, o := range q.outputs {
		if o.ImageID == imageID && o.DeletedAt == nil {
			results = append(results, o)
		}
	}

	return results, nil
}

func (q *MockOutputQuerier) UpdateOutputByID(ctx context.Context, confirmed bool, outputID, imageID, userID uuid.UUID) (*output.Output, error) {
	o, err := q.OutputByID(ctx, outputID, imageID, userID)
	if err != nil {
		return nil, err
	}

	o.Confirmed = confirmed
	o.UpdatedAt = time.Now().UTC()

	return o, nil
}

func (q *MockOutputQuerier) DeleteOutputByID(ctx context.Context, outputID, imageID, userID uuid.UUID) error {
	o, err := q.OutputByID(ctx, outputID, imageID, userID)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	o.DeletedAt = &now

	return nil
}

func (q *MockOutputQuerier) DeleteOutputsByImageID(ctx context.Context, imageID, userID uuid.UUID) error {
	if owner, ok := q.images[imageID]; !ok || owner != userID {
		return errors.New("unauthorized or missing image")
	}

	now := time.Now().UTC()
	for _, o := range q.outputs {
		if o.ImageID == imageID && o.DeletedAt == nil {
			o.DeletedAt = &now
		}
	}

	return nil
}

func (q *MockOutputQuerier) DeleteUserOutputs(ctx context.Context, userID uuid.UUID) error {
	found := false
	now := time.Now().UTC()
	for _, o := range q.outputs {
		if owner, ok := q.images[o.ImageID]; ok && owner == userID && o.DeletedAt == nil {
			found = true
			o.DeletedAt = &now
		}
	}

	if !found {
		return errors.New("not found")
	}

	return nil
}

func (m *MockOutputQuerier) Seed(outputs []*output.Output, imageOwnership map[uuid.UUID]uuid.UUID) {
	for _, o := range outputs {
		m.outputs[o.ID] = o
	}

	for imageID, userID := range imageOwnership {
		m.images[imageID] = userID
	}
}
