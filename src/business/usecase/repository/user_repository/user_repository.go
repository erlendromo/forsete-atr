package userrepository

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/user"
	"github.com/erlendromo/forsete-atr/src/database"
	"github.com/erlendromo/forsete-atr/src/querier"
	"github.com/erlendromo/forsete-atr/src/querier/sqlx"
	"github.com/google/uuid"
)

type UserRepository struct {
	querier querier.Querier[user.User]
}

func NewUserRepository(db database.Database) *UserRepository {
	return &UserRepository{
		querier: sqlx.NewSqlxQuerier[user.User](db),
	}
}

func (u *UserRepository) RegisterUser(ctx context.Context, email, hashedPassword string) (*user.User, error) {
	query := `
		WITH updated AS (
		  	UPDATE
				"user"
			SET
				password = $2,
			    deleted_at = NULL
			WHERE
				email = $1
			AND
				deleted_at IS NOT NULL
			RETURNING
				id
		)
		INSERT INTO
			"user" (email, password)
		SELECT
			$1,
			$2
		WHERE NOT EXISTS (
			SELECT
				1
			FROM
				updated
		)
		RETURNING
			id
	`

	return u.querier.QueryRowx(ctx, query, email, hashedPassword)
}

func (u *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	query := `
		SELECT
		    u.id,
		    u.email,
		    u.password,
			u.created_at,
			u.deleted_at,
		    r.name AS role_name
		FROM
		    "user" u
		LEFT JOIN
		    "role" r ON u.role_id = r.id
		WHERE
			u.id = $1
		AND
			u.deleted_at IS NULL
	`

	return u.querier.QueryRowx(ctx, query, id)
}

func (u *UserRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	query := `
		SELECT
		    u.id,
		    u.email,
		    u.password,
			u.created_at,
			u.deleted_at,
		    r.name AS role_name
		FROM
		    "user" u
		LEFT JOIN
		    "role" r ON u.role_id = r.id
		WHERE
			u.email = $1
		AND
			u.deleted_at IS NULL
	`

	return u.querier.QueryRowx(ctx, query, email)
}

func (u *UserRepository) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	query := `
		UPDATE
			"user"
		SET
			deleted_at = now()
		WHERE
			id = $1
		AND
			deleted_at IS NULL
	`

	return u.querier.Executex(ctx, query, id)
}
