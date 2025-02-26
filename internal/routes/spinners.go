package routes

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/matchstickn/sqlctest/assets/db"
	"github.com/matchstickn/sqlctest/internal/server"
)

func GetUserHandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var id server.UserId
		if err := c.BodyParser(&id); err != nil {
			return server.PublicWrapError(err, "bodyparser")
		}

		spinner, err := query.GetSpinner(ctx, id.Id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("spinner not found")
			}
			return server.PublicWrapError(err, "get")
		}
		return c.JSON(spinner)
	}
}

func GetUserTricksHandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var id server.UserId
		if err := c.BodyParser(&id); err != nil {
			return server.PublicWrapError(err, "bodyparser")
		}

		tricks, err := query.GetSpinnerTricks(ctx, id.Id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("tricks not found")
			}
			return server.PublicWrapError(err, "get")
		}
		return c.JSON(tricks)
	}
}
