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

func GetTrickHandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var id server.TrickId
		if err := c.BodyParser(&id); err != nil {
			return server.PublicWrapError(err, "bodyparser")
		}

		trick, err := query.GetTrick(ctx, id.Id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("trick not found")
			}
			return server.PublicWrapError(err, "get")
		}
		return c.JSON(trick)
	}
}

func ListTrickhandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		allTricks, err := query.GetAllTricks(ctx)
		if err != nil {
			return server.PublicWrapError(err, "list")
		}

		return c.JSON(allTricks)
	}
}

func CreateTrickHandler(ctx context.Context, query *db.Queries) fiber.Handler {
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

func DeleteTrickHandler(ctx context.Context, query *db.Queries) fiber.Handler {
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

func UpdateTrickHandler(ctx context.Context, query *db.Queries) fiber.Handler {
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
