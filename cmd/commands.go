package cmd

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jackc/pgx/v5"
	"github.com/matchstickn/sqlctest/assets/db"
	"github.com/matchstickn/sqlctest/internal/routes"
)

func SetUpRoutes(ctx context.Context, query *db.Queries, app *fiber.App) {
	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	// Tricks
	app.Get("/get", routes.GetTrickHandler(ctx, query))
	app.Get("/list", routes.ListTrickhandler(ctx, query))
	app.Post("/create", routes.CreateTrickHandler(ctx, query))
	app.Delete("/delete", routes.DeleteTrickHandler(ctx, query))
	app.Put("/update", routes.UpdateTrickHandler(ctx, query))
	// Auth
	if err := routes.SetUpAuthenticationHandlers(app); err != nil {
		log.Fatal(err)
	}

}

func SetUpDB(ctx context.Context, connstr string) (*db.Queries, *pgx.Conn) {
	pq, err := pgx.Connect(ctx, connstr)
	if err != nil {
		log.Fatal(err)
	}

	return db.New(pq), pq
}
