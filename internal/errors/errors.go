package errors

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type CustomErr struct {
	Message    string
	HttpStatus int
}

func (e *CustomErr) Error() string {
	return e.Message
}

var ErrorHandler = func(c *fiber.Ctx, err error) error {
	code := http.StatusInternalServerError

	var e *CustomErr
	if errors.As(err, &e) {
		code = e.HttpStatus
	}

	if err != nil {
		log.Errorf("Error: %w", err.Error())
		return c.Status(code).SendString(err.Error())
	}
	return nil
}
