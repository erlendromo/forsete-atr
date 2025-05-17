package output

import (
	"context"
	"testing"
	"time"

	"github.com/erlendromo/forsete-atr/src/business/domain/output"
	querier "github.com/erlendromo/forsete-atr/src/business/usecase/querier/output"
	"github.com/google/uuid"
)

type testCase struct {
	testName     string
	expectedPass bool

	id        uuid.UUID
	name      string
	format    string
	path      string
	imageID   uuid.UUID
	userID    uuid.UUID
	confirmed bool
}

func setup() *OutputRepository {
	now := time.Now().UTC()

	imageOneID := uuid.MustParse("f19cde2f-1c3e-4b63-96e9-dfa9b2f421cb")
	imageTwoID := uuid.MustParse("f19cde2f-1c3e-4b63-96e9-dfa9b2f421cc")
	userID := uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51462")

	populateOutputs := []*output.Output{
		{
			ID:        uuid.MustParse("11c8a0be-9298-4870-bc18-cddb47253d3b"),
			Name:      "output1",
			Format:    "json",
			Path:      "path/1",
			CreatedAt: now,
			ImageID:   imageOneID,
		},
		{
			ID:        uuid.MustParse("44f73aa6-f3e0-4de2-8b93-d4641efb8e84"),
			Name:      "deleted",
			Format:    "json",
			Path:      "path/2",
			CreatedAt: now,
			DeletedAt: &now,
			ImageID:   imageTwoID,
		},
	}

	populateOwnership := map[uuid.UUID]uuid.UUID{
		imageOneID: userID,
		imageTwoID: userID,
	}

	mockQuerier := querier.NewMockOutputQuerier()
	mockQuerier.Seed(populateOutputs, populateOwnership)

	return NewOutputRepository(mockQuerier)
}

func TestOutputRepository(t *testing.T) {
	t.Run("Register output", testRegisterOutput)
	t.Run("Get output by ID", testGetOutputByID)
	t.Run("Get outputs by image", testGetOutputsByImageID)
	t.Run("Update output", testUpdateOutputByID)
	t.Run("Delete output", testDeleteOutputByID)
	t.Run("Delete outputs by image", testDeleteOutputsByImageID)
	t.Run("Delete user outputs", testDeleteUserOutputs)
}

func testRegisterOutput(t *testing.T) {
	repo := setup()

	testCases := []testCase{
		{
			testName:     "Register valid output",
			expectedPass: true,

			name:    "newoutput",
			format:  "json",
			path:    "path/new",
			imageID: uuid.MustParse("f19cde2f-1c3e-4b63-96e9-dfa9b2f421cb"),
			userID:  uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51462"),
		},
		{
			testName:     "Register output invalid imageID",
			expectedPass: false,

			name:    "unauthorized",
			format:  "json",
			path:    "bad/path",
			imageID: uuid.New(),
			userID:  uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51462"),
		},
		{
			testName:     "Register output invalid userID",
			expectedPass: false,

			name:    "unauthorized",
			format:  "json",
			path:    "bad/path",
			imageID: uuid.MustParse("f19cde2f-1c3e-4b63-96e9-dfa9b2f421cc"),
			userID:  uuid.New(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := repo.RegisterOutput(context.Background(), tc.name, tc.format, tc.path, tc.imageID, tc.userID)
			if (err == nil) != tc.expectedPass {
				t.Errorf("unexpected result: %v", err)
			}
		})
	}
}

func testGetOutputByID(t *testing.T) {
	repo := setup()

	testCases := []testCase{
		{
			testName:     "Valid output",
			expectedPass: true,

			id:      uuid.MustParse("11c8a0be-9298-4870-bc18-cddb47253d3b"),
			imageID: uuid.MustParse("f19cde2f-1c3e-4b63-96e9-dfa9b2f421cb"),
			userID:  uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51462"),
		},
		{
			testName:     "Deleted output",
			expectedPass: false,

			id:      uuid.MustParse("44f73aa6-f3e0-4de2-8b93-d4641efb8e84"),
			imageID: uuid.MustParse("f19cde2f-1c3e-4b63-96e9-dfa9b2f421cc"),
			userID:  uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51462"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := repo.OutputByID(context.Background(), tc.id, tc.imageID, tc.userID)
			if (err == nil) != tc.expectedPass {
				t.Errorf("unexpected result: %v", err)
			}
		})
	}
}

func testGetOutputsByImageID(t *testing.T) {
	repo := setup()

	testCases := []testCase{
		{
			testName:     "Valid image",
			expectedPass: true,

			imageID: uuid.MustParse("f19cde2f-1c3e-4b63-96e9-dfa9b2f421cb"),
			userID:  uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51462"),
		},
		{
			testName:     "Invalid imageID",
			expectedPass: false,

			imageID: uuid.New(),
			userID:  uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51462"),
		},
		{
			testName:     "Invalid userID",
			expectedPass: false,

			imageID: uuid.MustParse("f19cde2f-1c3e-4b63-96e9-dfa9b2f421cc"),
			userID:  uuid.New(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := repo.OutputsByImageID(context.Background(), tc.imageID, tc.userID)
			if (err == nil) != tc.expectedPass {
				t.Errorf("unexpected result: %v", err)
			}
		})
	}
}

func testUpdateOutputByID(t *testing.T) {
	repo := setup()

	testCases := []testCase{
		{
			testName:     "Update valid output",
			expectedPass: true,

			id:        uuid.MustParse("11c8a0be-9298-4870-bc18-cddb47253d3b"),
			imageID:   uuid.MustParse("f19cde2f-1c3e-4b63-96e9-dfa9b2f421cb"),
			userID:    uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51462"),
			confirmed: true,
		},
		{
			testName:     "Fail to update non-existing output",
			expectedPass: false,

			id:        uuid.New(),
			imageID:   uuid.New(),
			userID:    uuid.New(),
			confirmed: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := repo.UpdateOutputByID(context.Background(), tc.confirmed, tc.id, tc.imageID, tc.userID)
			if (err == nil) != tc.expectedPass {
				t.Errorf("unexpected result: %v", err)
			}
		})
	}
}

func testDeleteOutputByID(t *testing.T) {
	repo := setup()

	testCases := []testCase{
		{
			testName:     "Delete valid output",
			expectedPass: true,

			id:      uuid.MustParse("11c8a0be-9298-4870-bc18-cddb47253d3b"),
			imageID: uuid.MustParse("f19cde2f-1c3e-4b63-96e9-dfa9b2f421cb"),
			userID:  uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51462"),
		},
		{
			testName:     "Delete invalid output",
			expectedPass: false,

			id:      uuid.New(),
			imageID: uuid.New(),
			userID:  uuid.New(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			err := repo.DeleteOutputByID(context.Background(), tc.id, tc.imageID, tc.userID)
			if (err == nil) != tc.expectedPass {
				t.Errorf("unexpected result: %v", err)
			}
		})
	}
}

func testDeleteOutputsByImageID(t *testing.T) {
	repo := setup()

	testCases := []testCase{
		{
			testName:     "Delete outputs by valid image",
			expectedPass: true,

			imageID: uuid.MustParse("f19cde2f-1c3e-4b63-96e9-dfa9b2f421cb"),
			userID:  uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51462"),
		},
		{
			testName:     "Delete outputs by invalid image",
			expectedPass: false,

			imageID: uuid.New(),
			userID:  uuid.New(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			err := repo.DeleteOutputsByImageID(context.Background(), tc.imageID, tc.userID)
			if (err == nil) != tc.expectedPass {
				t.Errorf("unexpected result: %v", err)
			}
		})
	}
}

func testDeleteUserOutputs(t *testing.T) {
	repo := setup()

	testCases := []testCase{
		{
			testName:     "Delete all outputs for user",
			expectedPass: true,

			userID: uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51462"),
		},
		{
			testName:     "Delete all outputs invalid user",
			expectedPass: false,

			userID: uuid.New(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			err := repo.DeleteUserOutputs(context.Background(), tc.userID)
			if (err == nil) != tc.expectedPass {
				t.Errorf("unexpected result: %v", err)
			}
		})
	}
}
