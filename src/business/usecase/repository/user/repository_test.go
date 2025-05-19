package userrepository

import (
	"context"
	"testing"
	"time"

	"github.com/erlendromo/forsete-atr/src/business/domain/user"
	querier "github.com/erlendromo/forsete-atr/src/business/usecase/querier/user"
	"github.com/erlendromo/forsete-atr/src/util"
	"github.com/google/uuid"
)

type testCase struct {
	testName     string
	expectedPass bool

	id        uuid.UUID
	email     string
	password  string
	createdAt time.Time
	deletedAt *time.Time
	roleID    int
	roleName  string
}

func setup() *UserRepository {
	now := time.Now().UTC()
	populateUsers := []*user.User{
		{
			ID:        uuid.MustParse("33219c9d-b3ef-43ce-b702-b52601e0f834"),
			Email:     "already@registered.com",
			Password:  "AlreadyRegisteredPassword123",
			CreatedAt: now,
			DeletedAt: nil,
			RoleID:    2,
			RoleName:  "Default",
		},
		{
			ID:        uuid.MustParse("33219c9d-b3ef-43ce-b702-b52601e0f835"),
			Email:     "deleted@user.com",
			Password:  "DeletedUserPassword123",
			CreatedAt: now,
			DeletedAt: &now,
			RoleID:    2,
			RoleName:  "Default",
		},
	}

	mockQuerier := querier.NewMockUserQuerier()
	mockQuerier.Seed(populateUsers)

	return NewUserRepository(mockQuerier)
}

func TestUserRepository(t *testing.T) {
	t.Run("Register user test", testRegisterUser)
	t.Run("Get by ID test", testGetByID)
	t.Run("Get by email test", testGetByEmail)
	t.Run("Delete by ID test", testDeleteByID)
}

func testRegisterUser(t *testing.T) {
	testUserRepo := setup()

	testCases := []testCase{
		{
			testName:     "Register new user",
			expectedPass: true,

			email:    "test@email.com",
			password: "SecretTestPassword123",
		},
		{
			testName:     "Email already registered",
			expectedPass: false,

			email:    "already@registered.com",
			password: "Em4i1AlreadyInUseSo5h0u1dF4i1",
		},
		{
			testName:     "Reactivate old user",
			expectedPass: true,

			email:    "deleted@user.com",
			password: "SecretTestPassword123",
		},
	}

	for _, tc := range testCases {

		t.Run(tc.testName, func(t *testing.T) {
			hashedPassword, errHash := util.HashPassword(tc.password)
			if errHash != nil {
				t.Errorf("unexpected error hashing password: %s", errHash.Error())
			}

			_, err := testUserRepo.RegisterUser(context.Background(), tc.email, hashedPassword)
			if (err == nil) != tc.expectedPass {
				if err != nil {
					t.Errorf("unexpected error when registering user: %s", err.Error())
				} else {
					t.Errorf("expected an error but got none")
				}
			}
		})

	}
}

func testGetByID(t *testing.T) {
	testUserRepo := setup()

	testCases := []testCase{
		{
			testName:     "Get valid user",
			expectedPass: true,

			id: uuid.MustParse("33219c9d-b3ef-43ce-b702-b52601e0f834"),
		},
		{
			testName:     "Get invalid user",
			expectedPass: false,

			id: uuid.New(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := testUserRepo.GetUserByID(context.Background(), tc.id)
			if (err == nil) != tc.expectedPass {
				if err != nil {
					t.Errorf("unexpected error when retrieving user by id: %s", err.Error())
				} else {
					t.Errorf("expected an error but got none")
				}
			}
		})
	}
}

func testGetByEmail(t *testing.T) {
	testUserRepo := setup()

	testCases := []testCase{
		{
			testName:     "Get valid user",
			expectedPass: true,

			email: "already@registered.com",
		},
		{
			testName:     "Get invalid user",
			expectedPass: false,

			email: "test@email.com",
		},
	}

	for _, tc := range testCases {

		t.Run(tc.testName, func(t *testing.T) {
			_, err := testUserRepo.GetUserByEmail(context.Background(), tc.email)
			if (err == nil) != tc.expectedPass {
				if err != nil {
					t.Errorf("unexpected error when retrieving user by email: %s", err.Error())
				} else {
					t.Errorf("expected an error but got none")
				}
			}
		})
	}
}

func testDeleteByID(t *testing.T) {
	testUserRepo := setup()

	testCases := []testCase{
		{
			testName:     "Delete valid user",
			expectedPass: true,

			id: uuid.MustParse("33219c9d-b3ef-43ce-b702-b52601e0f834"),
		},
		{
			testName:     "Delete invalid user",
			expectedPass: false,

			id: uuid.New(),
		},
	}

	for _, tc := range testCases {

		t.Run(tc.testName, func(t *testing.T) {
			err := testUserRepo.DeleteUserByID(context.Background(), tc.id)
			if (err == nil) != tc.expectedPass {
				if err != nil {
					t.Errorf("unexpected error when deleting user by id: %s", err.Error())
				} else {
					t.Errorf("expected an error but got none")
				}
			}
		})
	}
}
