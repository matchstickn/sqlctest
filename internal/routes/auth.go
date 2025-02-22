package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matchstickn/sqlctest/internal/auth"
)

func SetUpAuthenticationHandlers(app *fiber.App) error {
	if err := auth.GothSetUpRoutes(app); err != nil {
		return err
	}
	return nil
}
