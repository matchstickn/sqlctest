package server

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/matchstickn/sqlctest/assets/db"
)

func GetTrickHandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var trickid struct{ id int32 }
		if err := c.BodyParser(&trickid); err != nil {
			return err
		}

		trick, err := query.GetTrick(ctx, trickid.id)
		if err != nil {
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

		newTrickParams, err := BodyToCreateTrick(newTrick)
		if err != nil {
			return err
		}
		// newTrickParams := db.CreateTrickParams{
		// 	Name:  newTrick.Name,
		// 	Style: newTrick.Style,
		// 	Power: newTrick.Power,
		// }

		trick, err := query.CreateTrick(ctx, newTrickParams)
		if err != nil {
			return err
		}

		return c.JSON(trick)
	}
}
