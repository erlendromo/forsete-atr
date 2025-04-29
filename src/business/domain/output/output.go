package output

import (
	"time"

	"github.com/google/uuid"
)

type Output struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Path      string    `db:"path" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	Confirmed bool      `db:"confirmed" json:"confirmed"`
	ImageID   uuid.UUID `db:"image_id" json:"image_id"`
}
