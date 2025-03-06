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

func GetSpinnerHandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var id server.SpinnerId
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

func ListSpinnerHandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		spinner, err := query.ListSpinners(ctx)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("spinner not found")
			}
			return server.PublicWrapError(err, "list")
		}
		return c.JSON(spinner)
	}
}

func GetSpinnerTricksHandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var id server.SpinnerId
		if err := c.BodyParser(&id); err != nil {
			return server.PublicWrapError(err, "bodyparser")
		}

		Spinners, err := query.GetSpinnerTricks(ctx, id.Id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("Spinners not found")
			}
			return server.PublicWrapError(err, "get")
		}
		return c.JSON(Spinners)
	}
}

func CreateSpinnerHandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var newSpinner db.Spinner
		if err := c.BodyParser(&newSpinner); err != nil {
			return server.PublicWrapError(err, "bodyparser")
		}

		spinner, err := query.CreateSpinner(ctx, db.CreateSpinnerParams{
			Name:              newSpinner.Name,
			Email:             newSpinner.Email,
			Provider:          newSpinner.Provider,
			Tricks:            newSpinner.Tricks,
			Expiresat:         newSpinner.Expiresat,
			Accesstoken:       newSpinner.Accesstoken,
			Accesstokensecret: newSpinner.Accesstokensecret,
			Refreshtoken:      newSpinner.Refreshtoken,
		})
		if err != nil {
			return server.PublicWrapError(err, "create")
		}

		c.JSON(spinner)
		return c.SendStatus(fiber.StatusAccepted)
	}
}

func DeleteSpinnerHandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var id server.SpinnerId
		if err := c.BodyParser(&id); err != nil {
			return server.PublicWrapError(err, "bodyparser")
		}

		spinner, err := query.DeleteSpinner(ctx, id.Id)
		if err != nil {
			return server.PublicWrapError(err, "delete")
		}

		stringID := strconv.Itoa(int(id.Id))

		c.JSON(spinner)
		c.WriteString("Successfully Deleted Spinner with id: " + stringID)
		return c.SendStatus(fiber.StatusAccepted)
	}
}

func UpdateSpinnerHandler(ctx context.Context, query *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var newSpinner db.Spinner
		if err := c.BodyParser(&newSpinner); err != nil {
			return server.PublicWrapError(err, "bodyparser")
		}

		spinner, err := query.UpdateSpinner(ctx, db.UpdateSpinnerParams(newSpinner))
		if err != nil {
			return server.PublicWrapError(err, "update")
		}

		c.JSON(spinner)
		return c.SendStatus(fiber.StatusAccepted)
	}
}
