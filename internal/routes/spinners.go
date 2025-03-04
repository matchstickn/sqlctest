package routes

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

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

func CreateSpinnerHandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var newTrick db.Trick
		if err := c.BodyParser(&newTrick); err != nil {
			return server.PublicWrapError(err, "bodyparser")
		}

		trick, err := query.CreateTrick(ctx, db.CreateTrickParams{
			Name:  newTrick.Name,
			Style: newTrick.Style,
			Power: newTrick.Power,
		})
		if err != nil {
			return server.PublicWrapError(err, "create")
		}

		c.JSON(trick)
		return c.SendStatus(fiber.StatusAccepted)
	}
}

func DeleteSpinnerHandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var id server.TrickId
		if err := c.BodyParser(&id); err != nil {
			return server.PublicWrapError(err, "bodyparser")
		}

		if err := query.DeleteTrick(ctx, id.Id); err != nil {
			return server.PublicWrapError(err, "delete")
		}

		stringID := strconv.Itoa(int(id.Id))

		c.WriteString("Successfully Deleted trick with id: " + stringID)
		return c.SendStatus(fiber.StatusAccepted)
	}
}

func UpdateSpinnnerHandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var newTrick db.Trick
		if err := c.BodyParser(&newTrick); err != nil {
			return server.PublicWrapError(err, "bodyparser")
		}

		trick, err := query.UpdateTrick(ctx, db.UpdateTrickParams(newTrick))
		if err != nil {
			return server.PublicWrapError(err, "update")
		}

		c.JSON(trick)
		return c.SendStatus(fiber.StatusAccepted)
	}
}
