package routers

import (
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	Uri               string
	Method            string
	Function          func(c *fiber.Ctx) error
	HasAuthentication bool
}

// Registra as rotas http
func Register(app *fiber.App, routes *[]Route) {
	for _, route := range *routes {
		if route.HasAuthentication {
			// Adicionar o middleware de autenticacao
			app.Add(route.Method, route.Uri, route.Function)
		} else {
			app.Add(route.Method, route.Uri)
		}
	}
}
