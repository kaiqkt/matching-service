package routers

import (
	"matching-service/internal/controllers"
)

type RideRoutes struct {
	controller *controllers.RideController
}

func NewRideRoutes(c *controllers.RideController) *RideRoutes {
	return &RideRoutes{c}
}

func (rr *RideRoutes) Make() *[]Route {
	controller := rr.controller
	routes := []Route{
		{
			Uri:               "/ride-request",
			Method:            "GET",
			Function:          controller.RideRequest,
			HasAuthentication: true,
		},
	}

	return &routes
}
