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
	app.Route("/trick", func(api fiber.Router) {
		api.Get("/get", routes.GetTrickHandler(ctx, query))
		api.Get("/list", routes.ListTrickhandler(ctx, query))
		api.Post("/create", routes.CreateTrickHandler(ctx, query))
		api.Delete("/delete", routes.DeleteTrickHandler(ctx, query))
		api.Put("/update", routes.UpdateTrickHandler(ctx, query))
	}, "trick")

	// Spinners
	app.Route("/spinner", func(api fiber.Router) {
		api.Get("/get", routes.GetUserHandler(ctx, query))
		api.Get("/tricks", routes.GetUserTricksHandler(ctx, query))
	}, "spinner")
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
