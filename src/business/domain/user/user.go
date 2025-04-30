package user

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	Email     string     `db:"email" json:"email,omitempty"`
	Password  string     `db:"password" json:"-"`
	CreatedAt time.Time  `db:"created_at" json:"-"`
	DeletedAt *time.Time `db:"deleted_at" json:"-"`
	RoleID    int        `db:"role_id" json:"-"`
	RoleName  string     `db:"role_name" json:"-"`
}

// If err is nil -> password is correct
func (u *User) ComparePassword(unhashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(unhashedPassword))
}

func (u *User) IsAdmin() bool {
	return u.RoleName == "Admin"
}
