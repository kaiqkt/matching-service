package main

import (
	"fmt"
	"matching-service/cmd/api/di"
	cerrors "matching-service/internal/errors"
	"matching-service/internal/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: cerrors.ErrorHandler,
	})
	app.Use(logger.New())
	serverPort := di.GetEnv("SERVER_PORT")

	db, err := di.NewPostgresDB()

	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	ridersRoutes := di.BuildRideDI(db)

	routes := appendRoutes(ridersRoutes)
	routers.Register(app, routes)

	app.Listen(fmt.Sprintf(":%s", serverPort))
}

func appendRoutes(routes ...*[]routers.Route) *[]routers.Route {
	var result []routers.Route
	for _, route := range routes {
		result = append(result, *route...)
	}

	return &result
}
