package session

import (
	"context"
	"testing"
	"time"

	"github.com/erlendromo/forsete-atr/src/business/domain/session"
	querier "github.com/erlendromo/forsete-atr/src/business/usecase/querier/session"
	"github.com/google/uuid"
)

type testCase struct {
	testName     string
	expectedPass bool

	token     uuid.UUID
	userID    uuid.UUID
	createdAt time.Time
	expiresAt time.Time
}

func setup() *SessionRepository {
	now := time.Now().UTC()
	populateSessions := []*session.Session{
		{
			Token:     uuid.MustParse("c59f1ac5-8a28-4031-8d8d-d3d20d40796a"),
			UserID:    uuid.MustParse("c59f1ac5-8a28-4031-8d8d-d3d20d40796a"),
			CreatedAt: now,
			ExpiresAt: now.Add(time.Hour * 1),
		},
		{
			Token:     uuid.MustParse("c59f1ac5-8a28-4031-8d8d-d3d20d40796b"),
			UserID:    uuid.MustParse("c59f1ac5-8a28-4031-8d8d-d3d20d40796b"),
			CreatedAt: now,
			ExpiresAt: now.Add(-time.Hour * 1),
		},
	}

	querier := querier.NewMockSessionQuerier()
	querier.Seed(populateSessions)

	return NewSessionRepository(querier)
}

func TestSessionRepository(t *testing.T) {
	t.Run("Create session test", testCreateSession)
	t.Run("Get valid session test", testGetValidSession)
	t.Run("Delete session test", testDeleteSession)
	t.Run("Clear sessions by user id test", testClearSessionsByUserID)
	t.Run("Clear expired sessions test", testClearExpiredSessions)
}

func testCreateSession(t *testing.T) {
	testSessionsRepo := setup()

	now := time.Now().UTC()
	testCases := []testCase{
		{
			testName:     "Create valid session",
			expectedPass: true,

			token:     uuid.New(),
			userID:    uuid.New(),
			createdAt: now,
			expiresAt: now.Add(time.Hour * 1),
		},
	}

	for _, tc := range testCases {

		t.Run(tc.testName, func(t *testing.T) {
			_, err := testSessionsRepo.CreateSession(context.Background(), tc.token)
			if (err == nil) != tc.expectedPass {
				if err != nil {
					t.Errorf("unexpected error when creating session: %s", err.Error())
				} else {
					t.Errorf("expected an error but got none")
				}
			}
		})

	}
}

func testGetValidSession(t *testing.T) {
	testSessionsRepo := setup()

	testCases := []testCase{
		{
			testName:     "Get valid session",
			expectedPass: true,

			token: uuid.MustParse("c59f1ac5-8a28-4031-8d8d-d3d20d40796a"),
		},
		{
			testName:     "Get expired session",
			expectedPass: false,

			token: uuid.MustParse("c59f1ac5-8a28-4031-8d8d-d3d20d40796b"),
		},
		{
			testName:     "Get invalid session",
			expectedPass: false,

			token: uuid.New(),
		},
	}

	for _, tc := range testCases {

		t.Run(tc.testName, func(t *testing.T) {
			_, err := testSessionsRepo.GetValidSession(context.Background(), tc.token)
			if (err == nil) != tc.expectedPass {
				if err != nil {
					t.Errorf("unexpected error when creating session: %s", err.Error())
				} else {
					t.Errorf("expected an error but got none")
				}
			}
		})

	}
}

func testDeleteSession(t *testing.T) {
	testSessionsRepo := setup()

	testCases := []testCase{
		{
			testName:     "Delete valid session",
			expectedPass: true,

			token:  uuid.MustParse("c59f1ac5-8a28-4031-8d8d-d3d20d40796a"),
			userID: uuid.MustParse("c59f1ac5-8a28-4031-8d8d-d3d20d40796a"),
		},
		{
			testName:     "Delete invalid session",
			expectedPass: false,

			token:  uuid.New(),
			userID: uuid.New(),
		},
	}

	for _, tc := range testCases {

		t.Run(tc.testName, func(t *testing.T) {
			err := testSessionsRepo.DeleteSession(context.Background(), tc.token, tc.userID)
			if (err == nil) != tc.expectedPass {
				if err != nil {
					t.Errorf("unexpected error when creating session: %s", err.Error())
				} else {
					t.Errorf("expected an error but got none")
				}
			}
		})

	}
}

func testClearSessionsByUserID(t *testing.T) {
	testSessionsRepo := setup()

	testCases := []testCase{
		{
			testName:     "Clear valid user sessions",
			expectedPass: true,

			userID: uuid.MustParse("c59f1ac5-8a28-4031-8d8d-d3d20d40796a"),
		},
		{
			testName:     "Clear invalid user sessions",
			expectedPass: false,

			userID: uuid.New(),
		},
	}

	for _, tc := range testCases {

		t.Run(tc.testName, func(t *testing.T) {
			err := testSessionsRepo.ClearSessionsByUserID(context.Background(), tc.userID)
			if (err == nil) != tc.expectedPass {
				if err != nil {
					t.Errorf("unexpected error when creating session: %s", err.Error())
				} else {
					t.Errorf("expected an error but got none")
				}
			}
		})

	}
}

func testClearExpiredSessions(t *testing.T) {
	testSessionsRepo := setup()

	testCases := []testCase{
		{
			testName:     "Clear expired sessions",
			expectedPass: true,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.testName, func(t *testing.T) {
			err := testSessionsRepo.ClearExpiredSessions(context.Background())
			if (err == nil) != tc.expectedPass {
				if err != nil {
					t.Errorf("unexpected error when creating session: %s", err.Error())
				} else {
					t.Errorf("expected an error but got none")
				}
			}
		})

	}
}
