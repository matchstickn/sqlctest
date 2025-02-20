package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/matchstickn/sqlctest/assets/db"
	"github.com/matchstickn/sqlctest/internal/auth"
)

func LogoutHandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return auth.GothLogout(c)
	}
}

func AuthenticationHandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return auth.GothAuthenticate(c)
	}
}

func AuthenticationCallbackHandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return auth.GothAuthenitcationCallback(c)
	}
}
