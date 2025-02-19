package main

import (
	"context"
	"log"
	"os"

	"github.com/matchstickn/sqlctest/cmd"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	if err := godotenv.Load(); err != nil {
		log.Println("detected env injection, if not:", err)
	}
	connstr := os.Getenv("NEONTECH_URL")

	query, pq := cmd.SetUpDB(ctx, connstr)
	defer pq.Close(ctx)

	app := fiber.New(fiber.Config{})

	cmd.SetUpRoutes(ctx, query, app)

	log.Fatal(app.Listen(":4000"))
}
