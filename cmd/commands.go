package cmd

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5"
	"github.com/matchstickn/sqlctest/assets/db"
	"github.com/matchstickn/sqlctest/internal/server"
)

func SetUpFiber(ctx context.Context, query *db.Queries, app *fiber.App) {
	app.Use(logger.New())
	app.Get("/get", server.GetTrickHandler(ctx, query))
	app.Get("/list", server.ListTrickhandler(ctx, query))
	app.Post("/create", server.CreateTrickHandler(ctx, query))
	app.Delete("/delete", server.DeleteTrickHandler(ctx, query))
	app.Put("/update", server.UpdateTrickHandler(ctx, query))
}

func SetUpDB(ctx context.Context, connstr string) *db.Queries {
	pq, err := pgx.Connect(ctx, connstr)
	if err != nil {
		log.Fatal(err)
	}
	defer pq.Close(ctx)

	return db.New(pq)
}
