package session

import (
	"time"

	"github.com/google/uuid"
)

// Session
//
//	@Summary		Session
//	@Description	Session containing token etc.
type Session struct {
	Token     uuid.UUID `db:"token" json:"token"`
	UserID    uuid.UUID `db:"user_id" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	ExpiresAt time.Time `db:"expires_at" json:"-"`
}
