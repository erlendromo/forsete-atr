package userrepository

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/business/domain/auth/user"
	"github.com/erlendromo/forsete-atr/src/database"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) RegisterUser(ctx context.Context, email, hashedPassword string) (*user.User, error) {
	query := `
	INSERT INTO
		users (email, password)
	VALUES
		($1, $2)
	ON CONFLICT
		(email)
	DO UPDATE SET
		email = EXCLUDED.email
	RETURNING
		id
	`

	return database.QueryRowx[user.User](ctx, u.db, query, email, hashedPassword)
}

func (u *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	query := `
	SELECT
	    u.id,
	    u.email,
	    u.password,
	    r.name AS role_name
	FROM
	    users u
	LEFT JOIN
	    roles r ON u.role_id = r.id
	WHERE
		u.id = $1
	`

	return database.QueryRowx[user.User](ctx, u.db, query, id)
}

func (u *UserRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	query := `
	SELECT
	    u.id,
	    u.email,
	    u.password,
	    r.name AS role_name
	FROM
	    users u
	LEFT JOIN
	    roles r ON u.role_id = r.id
	WHERE
		u.email = $1
	`

	return database.QueryRowx[user.User](ctx, u.db, query, email)
}
