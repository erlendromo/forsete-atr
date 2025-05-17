package image

import (
	"context"
	"testing"
	"time"

	"github.com/erlendromo/forsete-atr/src/business/domain/image"
	querier "github.com/erlendromo/forsete-atr/src/business/usecase/querier/image"
	"github.com/google/uuid"
)

type testCase struct {
	testName     string
	expectedPass bool

	id     uuid.UUID
	name   string
	format string
	path   string
	userID uuid.UUID
}

func setup() *ImageRepository {
	now := time.Now().UTC()

	populateImages := []*image.Image{
		{
			ID:         uuid.MustParse("11c8a0be-9298-4870-bc18-cddb47253d3b"),
			Name:       "image1",
			Format:     "jpg",
			Path:       "path/img1",
			UploadedAt: now,
			UserID:     uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51462"),
		},
		{
			ID:         uuid.MustParse("44f73aa6-f3e0-4de2-8b93-d4641efb8e84"),
			Name:       "deletedImage",
			Format:     "png",
			Path:       "path/img2",
			UploadedAt: now,
			DeletedAt:  &now,
			UserID:     uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51463"),
		},
	}

	mockQuerier := querier.NewMockImageQuerier()
	mockQuerier.Seed(populateImages)

	return NewImageRepository(mockQuerier)
}

func TestImageRepository(t *testing.T) {
	t.Run("Register image", testRegisterImage)
	t.Run("Get image by ID", testGetImageByID)
	t.Run("Get images by user ID", testGetImagesByUserID)
	t.Run("Delete image by ID", testDeleteImageByID)
	t.Run("Delete user images", testDeleteUserImages)
}

func testRegisterImage(t *testing.T) {
	repo := setup()

	testCases := []testCase{
		{
			testName:     "Register valid image 1",
			expectedPass: true,

			name:   "imgX",
			format: "jpg",
			path:   "path/new",
			userID: uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51462"),
		},
		{
			testName:     "Register valid image 2", // Assume userID is active user
			expectedPass: true,

			name:   "fail",
			format: "png",
			path:   "path/empty",
			userID: uuid.New(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := repo.RegisterImage(context.Background(), tc.name, tc.format, tc.path, tc.userID)
			if (err == nil) != tc.expectedPass {
				t.Errorf("unexpected result: %v", err)
			}
		})
	}
}

func testGetImageByID(t *testing.T) {
	repo := setup()

	testCases := []testCase{
		{
			testName:     "Valid image",
			expectedPass: true,

			id:     uuid.MustParse("11c8a0be-9298-4870-bc18-cddb47253d3b"),
			userID: uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51462"),
		},
		{
			testName:     "Deleted image",
			expectedPass: false,

			id:     uuid.MustParse("44f73aa6-f3e0-4de2-8b93-d4641efb8e84"),
			userID: uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51463"),
		},
		{
			testName:     "Invalid imageID",
			expectedPass: false,

			id:     uuid.New(),
			userID: uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51462"),
		},
		{
			testName:     "Invalid user",
			expectedPass: false,

			id:     uuid.MustParse("11c8a0be-9298-4870-bc18-cddb47253d3b"),
			userID: uuid.New(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := repo.ImageByID(context.Background(), tc.id, tc.userID)
			if (err == nil) != tc.expectedPass {
				t.Errorf("unexpected result: %v", err)
			}
		})
	}
}

func testGetImagesByUserID(t *testing.T) {
	repo := setup()

	testCases := []testCase{
		{
			testName:     "Valid user",
			expectedPass: true,

			userID: uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51462"),
		},
		{
			testName:     "Invalid user",
			expectedPass: false,

			userID: uuid.New(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := repo.ImagesByUserID(context.Background(), tc.userID)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func testDeleteImageByID(t *testing.T) {
	repo := setup()

	testCases := []testCase{
		{
			testName:     "Delete valid image",
			expectedPass: true,

			id:     uuid.MustParse("11c8a0be-9298-4870-bc18-cddb47253d3b"),
			userID: uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51462"),
		},
		{
			testName:     "Invalid imageID",
			expectedPass: false,

			id:     uuid.New(),
			userID: uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51462"),
		},
		{
			testName:     "Invalid userID",
			expectedPass: false,

			id:     uuid.MustParse("11c8a0be-9298-4870-bc18-cddb47253d3b"),
			userID: uuid.New(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			err := repo.DeleteImageByID(context.Background(), tc.id, tc.userID)
			if (err == nil) != tc.expectedPass {
				t.Errorf("unexpected result: %v", err)
			}
		})
	}
}

func testDeleteUserImages(t *testing.T) {
	repo := setup()

	testCases := []testCase{
		{
			testName:     "Delete images for valid user",
			expectedPass: true,

			userID: uuid.MustParse("bd98c8e7-093c-472b-8183-d1fd09f51462"),
		},
		{
			testName:     "Delete images for invalid user",
			expectedPass: false,

			userID: uuid.New(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			err := repo.DeleteUserImages(context.Background(), tc.userID)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
