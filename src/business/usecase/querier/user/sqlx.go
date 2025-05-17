package user

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/user"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SQLUserQuerier struct {
	db *sqlx.DB
}

func NewSQLUserQuerier(db *sqlx.DB) *SQLUserQuerier {
	return &SQLUserQuerier{
		db: db,
	}
}

func (q *SQLUserQuerier) RegisterUser(ctx context.Context, email, hashedPassword string) (*user.User, error) {
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

	var u user.User
	err := q.db.QueryRowxContext(ctx, query, email, hashedPassword).StructScan(&u)

	return &u, err
}

func (q *SQLUserQuerier) GetUserByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
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

	var u user.User
	err := q.db.QueryRowxContext(ctx, query, id).StructScan(&u)

	return &u, err
}

func (q *SQLUserQuerier) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
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

	var u user.User
	err := q.db.QueryRowxContext(ctx, query, email).StructScan(&u)

	return &u, err
}

func (q *SQLUserQuerier) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
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

	_, err := q.db.ExecContext(ctx, query, id)
	return err
}
