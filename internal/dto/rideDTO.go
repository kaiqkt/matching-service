package dto

import (
	"matching-service/internal/models"
	"time"

	"github.com/oklog/ulid/v2"
)

type RideDTO struct {
	ID          string         `json:"id"`
	Source      *CoordinateDTO `json:"source"`
	Destination *CoordinateDTO `json:"destination"`
	Status      string         `json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
}

func (r *RideDTO) ToDomain() *models.Ride {
	source := r.Source.toDomain()
	destination := r.Destination.toDomain()

	return &models.Ride{
		ID:          ulid.Make(),
		UserID:      ulid.Make().String(),
		Source:      source,
		Destination: destination,
		Status:      models.WAITING_DRIVER,
		CreatedAt:   time.Now(),
	}
}
