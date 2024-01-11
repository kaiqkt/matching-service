package services

import (
	"fmt"
	cerros "matching-service/internal/errors"
	"matching-service/internal/models"
	"matching-service/internal/repositories"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
)

type RideService struct {
	repository *repositories.RideRepository
}

func NewRideService(repository *repositories.RideRepository) *RideService {
	return &RideService{repository}
}

func (s *RideService) CreateRideRequest(ride *models.Ride) error {
	exists, err := s.repository.ExistsByUserID(&ride.UserID)
	if err != nil {
		return err
	}

	if exists {
		return &cerros.CustomErr{Message: fmt.Sprintf("Ride already exists for user id %s", ride.UserID), HttpStatus: http.StatusConflict}
	}

	if err = s.repository.Save(ride); err != nil {
		return err
	}

	//enviar para uma fila

	log.Infof("Created ride request %s", ride.ID)

	return nil
}
