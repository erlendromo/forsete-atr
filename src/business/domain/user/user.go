package user

import (
	"os"
	"path"
	"time"

	"github.com/erlendromo/forsete-atr/src/util"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User
//
//	@Summary		User
//	@Description	User containing id, email etc.
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

func (u *User) CreateDirs() error {
	basePath := path.Join(util.USERS_PATH, u.ID.String())

	imagesPath := path.Join(basePath, util.IMAGES)
	if err := os.MkdirAll(imagesPath, os.ModeDir); err != nil {
		return err
	}

	outputsPath := path.Join(basePath, util.OUTPUTS)
	if err := os.MkdirAll(outputsPath, os.ModeDir); err != nil {
		return err
	}

	return nil
}

func (u *User) RemoveData() error {
	return os.RemoveAll(path.Join(util.USERS_PATH, u.ID.String()))
}
