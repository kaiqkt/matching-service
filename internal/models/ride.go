package models

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Ride struct {
	ID          ulid.ULID   `json:"id"`
	UserID      string      `json:"user_id"`
	Source      *Coordinate `json:"source"`
	Destination *Coordinate `json:"destination"`
	Status      string      `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
}
