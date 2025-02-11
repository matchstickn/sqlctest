package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/matchstickn/sqlctest/assets/db"
)

func GetTrickHandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var id trickId
		if err := c.BodyParser(&id); err != nil {
			return err
		}

		fmt.Println(id.Id)

		trick, err := query.GetTrick(ctx, id.Id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("user not found")
			}
			return err
		}
		return c.JSON(trick)
	}
}

func ListTrickhandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		allTricks, err := query.GetAllTricks(ctx)
		if err != nil {
			return err
		}

		return c.JSON(allTricks)
	}
}

func CreateTrickHandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var newTrick db.Trick
		if err := c.BodyParser(&newTrick); err != nil {
			return err
		}
		fmt.Println(newTrick)

		trick, err := query.CreateTrick(ctx, db.CreateTrickParams{
			Name:  newTrick.Name,
			Style: newTrick.Style,
			Power: newTrick.Power,
		})

		if err != nil {
			return err
		}

		fmt.Println(trick)
		c.JSON(trick)
		return c.SendStatus(fiber.StatusAccepted)
	}
}

func DeleteTrickHandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var id trickId
		if err := c.BodyParser(&id); err != nil {
			return err
		}

		if err := query.DeleteTrick(ctx, id.Id); err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusAccepted)
	}
}

func UpdateTrickHandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var newTrick db.Trick
		if err := c.BodyParser(&newTrick); err != nil {
			return err
		}
		fmt.Println(newTrick)

		trick, err := query.UpdateTrick(ctx, db.UpdateTrickParams(newTrick))

		if err != nil {
			return err
		}

		c.JSON(trick)
		return c.SendStatus(fiber.StatusAccepted)
	}
}
