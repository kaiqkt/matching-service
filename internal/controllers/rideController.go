package controllers

import (
	"matching-service/internal/dto"
	"matching-service/internal/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type RideController struct {
	service *services.RideService
}

func NewRideController(rs *services.RideService) *RideController {
	return &RideController{rs}
}

func (controller *RideController) RideRequest(c *fiber.Ctx) error {
	rs := controller.service
	request := new(dto.RideDTO)

	err := c.BodyParser(request)
	if err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}

	ride := request.ToDomain()

	err = rs.CreateRideRequest(ride)
	if err != nil {
		panic(err)
	}

	c.JSON(ride)
	return c.SendStatus(http.StatusCreated)
}
