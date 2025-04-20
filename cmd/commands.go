package cmd

import (
	"context"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jackc/pgx/v5"
	"github.com/matchstickn/sqlctest/assets/db"
	"github.com/matchstickn/sqlctest/internal/auth"
	"github.com/matchstickn/sqlctest/internal/routes"
	"github.com/matchstickn/sqlctest/internal/server"
)

func SetUpRoutes(ctx context.Context, query *db.Queries, app *fiber.App, v *validator.Validate) {
	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cache.New())
	app.Use(helmet.New())
	app.Use(limiter.New(limiter.Config{
		Expiration: time.Minute * 2,
		Max:        8,
	}))
	// Auth:Open
	app.Route("/auth", func(au fiber.Router) {
		au.Post("/signup", routes.SignUp())
		au.Post("/authenticate", routes.GetAuthenticatedAccessToken())
		au.Get("/redirect", routes.SocialLoginRedirect())
		au.Post("/pass", routes.ChangePassword())
		// Use this or webstudio.is system.search
		au.Get("/callback", routes.SocialLoginCallback())
	})
	// Protected
	app.Use(auth.EnsureValidToken())
	app.Get("/validate", routes.Test)
	// Tricks
	app.Route("/trick", func(api fiber.Router) {
		api.Get("/get", routes.GetTrickHandler(ctx, query, v))
		api.Get("/list", routes.ListTrickhandler(ctx, query, v))
		api.Post("/create", routes.CreateTrickHandler(ctx, query, v))
		api.Delete("/delete", routes.DeleteTrickHandler(ctx, query, v))
		api.Put("/update", routes.UpdateTrickHandler(ctx, query, v))
	}, "trick")

	// Spinners
	app.Route("/spinner", func(api fiber.Router) {
		api.Get("/get", routes.GetSpinnerHandler(ctx, query))
		api.Get("/tricks", routes.GetSpinnerTricksHandler(ctx, query))
		api.Get("/list", routes.ListSpinnerHandler(ctx, query))
		api.Post("/create", routes.CreateSpinnerHandler(ctx, query))
		api.Delete("/delete", routes.DeleteSpinnerHandler(ctx, query))
		api.Put("/update", routes.UpdateSpinnerHandler(ctx, query))
	}, "spinner")

}

func SetUpDB(ctx context.Context, connstr string) (*db.Queries, *pgx.Conn) {
	pq, err := pgx.Connect(ctx, connstr)
	if err != nil {
		log.Fatal(err)
	}

	return db.New(pq), pq
}

func SetUpValidator() *validator.Validate {
	return server.NewValidator()
}

func RecoveryInputFunc(fn func(*bool)) {
	Recovered := new(bool)
	retries := 0

	for !*Recovered {
		fn(Recovered)
		if retries == 3 {
			break
		}
		retries++
	}
	log.Println("unrecoverable")
}
