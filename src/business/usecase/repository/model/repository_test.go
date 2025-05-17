package model

import (
	"context"
	"testing"

	"github.com/erlendromo/forsete-atr/src/business/domain/model"
	querier "github.com/erlendromo/forsete-atr/src/business/usecase/querier/model"
)

type testCase struct {
	testName     string
	expectedPass bool

	id          int
	name        string
	path        string
	modelTypeID int
	modelType   string
}

func setup() *ModelRepository {
	populateModels := []*model.Model{
		{
			ID:          1,
			Name:        "model1",
			Path:        "path/1",
			ModelTypeID: 1,
			ModelType:   "regionsegmentation",
		},
		{
			ID:          2,
			Name:        "model2",
			Path:        "path/2",
			ModelTypeID: 2,
			ModelType:   "linesegmentation",
		},
		{
			ID:          3,
			Name:        "model3",
			Path:        "path/3",
			ModelTypeID: 3,
			ModelType:   "textrecognition",
		},
	}

	mockQuerier := querier.NewMockModelQuerier(4)
	mockQuerier.Seed(populateModels)

	return NewModelRepository(mockQuerier)
}

func TestModelRepository(t *testing.T) {
	t.Run("Register model", testRegisterModel)
	t.Run("Get model by ID", testModelByID)
	t.Run("Get model by name", testModelByName)
	t.Run("Get models by type", testModelsByType)
	t.Run("Get all models", testAllModels)
}

func testRegisterModel(t *testing.T) {
	testModelRepo := setup()

	testCases := []testCase{
		{
			testName:     "Register valid model",
			expectedPass: true,

			name:        "newmodel",
			path:        "/models/newmodel",
			modelTypeID: 1,
		},
		{
			testName:     "Register model with invalid type ID",
			expectedPass: false,

			name:        "invalidmodel",
			path:        "/models/invalidmodel",
			modelTypeID: 99,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			model, err := testModelRepo.RegisterModel(context.Background(), tc.name, tc.path, tc.modelTypeID)
			if (err == nil) != tc.expectedPass {
				t.Errorf("unexpected result: %v", err)
			}

			if err == nil && model.Name != tc.name {
				t.Errorf("expected name %s, got %s", tc.name, model.Name)
			}
		})
	}
}

func testModelByID(t *testing.T) {
	testModelRepo := setup()

	testCases := []testCase{
		{
			testName:     "Get existing model by ID",
			expectedPass: true,

			id: 1,
		},
		{
			testName:     "Get non-existent model by ID",
			expectedPass: false,

			id: 999,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			model, err := testModelRepo.ModelByID(context.Background(), tc.id)
			if (err == nil) != tc.expectedPass {
				t.Errorf("unexpected result: %v", err)
			}

			if err == nil && model.ID != tc.id {
				t.Errorf("expected ID %d, got %d", tc.id, model.ID)
			}
		})
	}
}

func testModelByName(t *testing.T) {
	testModelRepo := setup()

	testCases := []testCase{
		{
			testName:     "Get existing model by name",
			expectedPass: true,

			name: "model1",
		},
		{
			testName:     "Get non-existent model by name",
			expectedPass: false,

			name: "unknown",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			model, err := testModelRepo.ModelByName(context.Background(), tc.name)
			if (err == nil) != tc.expectedPass {
				t.Errorf("unexpected result: %v", err)
			}

			if err == nil && model.Name != tc.name {
				t.Errorf("expected name %s, got %s", tc.name, model.Name)
			}
		})
	}
}

func testModelsByType(t *testing.T) {
	testModelRepo := setup()

	testCases := []testCase{
		{
			testName:     "Get models by valid type",
			expectedPass: true,

			modelType: "regionsegmentation",
		},
		{
			testName:     "Get models by invalid type",
			expectedPass: false,

			modelType: "invalidtype",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := testModelRepo.ModelsByType(context.Background(), tc.modelType)
			if (err == nil) != tc.expectedPass {
				t.Errorf("unexpected result: %v", err)
			}
		})
	}
}

func testAllModels(t *testing.T) {
	testModelRepo := setup()

	_, err := testModelRepo.AllModels(context.Background())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
