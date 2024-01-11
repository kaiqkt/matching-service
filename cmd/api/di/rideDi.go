package di

import (
	"database/sql"
	"matching-service/internal/controllers"
	"matching-service/internal/repositories"
	"matching-service/internal/routers"
	"matching-service/internal/services"
)

func BuildRideDI(db *sql.DB) *[]routers.Route {
	repository := repositories.NewRideRepository(db)
	service := services.NewRideService(repository)
	controller := controllers.NewRideController(service)
	routers := routers.NewRideRoutes(controller)

	return routers.Make()
}
