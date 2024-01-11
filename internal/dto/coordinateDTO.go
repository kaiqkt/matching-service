package dto

import "matching-service/internal/models"

type CoordinateDTO struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (c *CoordinateDTO) toDomain() *models.Coordinate {
	return &models.Coordinate{
		Latitude:  c.Latitude,
		Longitude: c.Longitude,
	}
}
