package pipeline

import (
	"context"
	"testing"

	"github.com/erlendromo/forsete-atr/src/business/domain/pipeline"
	querier "github.com/erlendromo/forsete-atr/src/business/usecase/querier/pipeline"
)

type testCase struct {
	testName     string
	expectedPass bool

	id         int
	name       string
	path       string
	pipelineID int
	modelID    int
}

func setup() *PipelineRepository {
	populatePipelines := []*pipeline.Pipeline{
		{
			ID:   1,
			Name: "pipeline1",
			Path: "path/1",
		},
		{
			ID:   2,
			Name: "pipeline2",
			Path: "path/2",
		},
	}

	populatePipelineModels := []map[int][]int{
		{
			1: {1, 2},
		},
		{
			2: {1, 3},
		},
	}

	mockQuerier := querier.NewMockPipelineQuerier(3)
	mockQuerier.Seed(populatePipelines, populatePipelineModels)

	return NewPipelineRepository(mockQuerier)
}

func TestPipelineRepository(t *testing.T) {
	t.Run("Register pipeline test", testRegisterPipeline)
	t.Run("Pipeline by ID test", testPipelineByID)
	t.Run("List pipelines test", testListPipelines)
	t.Run("Register pipeline model test", testRegisterPipelineModel)
}

func testRegisterPipeline(t *testing.T) {
	repo := setup()

	testCases := []testCase{
		{
			testName:     "Register new pipeline",
			expectedPass: true,

			name: "pipeline3",
			path: "path/3",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := repo.RegisterPipeline(context.Background(), tc.name, tc.path)
			if (err == nil) != tc.expectedPass {
				t.Fatalf("expected error? %v, got: %v", !tc.expectedPass, err)
			}
		})
	}
}

func testPipelineByID(t *testing.T) {
	repo := setup()

	testCases := []testCase{
		{
			testName:     "Valid pipeline by ID",
			expectedPass: true,

			id: 1,
		},
		{
			testName:     "Invalid pipeline by ID",
			expectedPass: false,

			id: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := repo.PipelineByID(context.Background(), tc.id)
			if (err == nil) != tc.expectedPass {
				t.Fatalf("expected error? %v, got: %v", !tc.expectedPass, err)
			}
		})
	}
}

func testListPipelines(t *testing.T) {
	repo := setup()

	t.Run("List all pipelines", func(t *testing.T) {
		_, err := repo.ListPipelines(context.Background())
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

func testRegisterPipelineModel(t *testing.T) {
	repo := setup()

	testCases := []testCase{
		{
			testName:     "Register valid pipeline model",
			expectedPass: true,

			pipelineID: 1,
			modelID:    99,
		},
		{
			testName:     "Register invalid pipeline model",
			expectedPass: false,

			pipelineID: 9999,
			modelID:    1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			err := repo.RegisterPipelineModel(context.Background(), tc.pipelineID, tc.modelID)
			if (err == nil) != tc.expectedPass {
				if err != nil {
					t.Errorf("expected error? %v, got: %v", !tc.expectedPass, err)
				} else {
					t.Error("expected error but got none")
				}
			}
		})
	}
}
