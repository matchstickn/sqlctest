package main

import (
	"context"
	"log"
	"os"

	"github.com/matchstickn/sqlctest/assets/db"
	"github.com/matchstickn/sqlctest/internal/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	if err := godotenv.Load(); err != nil {
		log.Println("detected env injection, if not:", err)
	}

	// connstr := "postgres", "user=postgres password="+os.Getenv("password")+" dbname=FlintCRUD sslmode=disable"
	connstr := os.Getenv("NEONTECH_URL")

	pq, err := pgx.Connect(ctx, connstr)
	if err != nil {
		log.Fatal(err)
	}
	defer pq.Close(ctx)

	query := db.New(pq)

	app := fiber.New(fiber.Config{})

	app.Use(logger.New())
	app.Get("/get", server.GetTrickHandler(ctx, query))
	app.Get("/list", server.ListTrickhandler(ctx, query))
	app.Post("/create", server.CreateTrickHandler(ctx, query))
	app.Delete("/delete", server.DeleteTrickHandler(ctx, query))
	app.Put("/update", server.UpdateTrickHandler(ctx, query))

	log.Fatal(app.Listen(":4000"))
}
