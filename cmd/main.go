package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/matchstickn/sqlctest/assets/db"
	"github.com/matchstickn/sqlctest/internal/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/internal/schema"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	if err := godotenv.Load(); err != nil {
		log.Println("detected env injection, if not:", err)
	}

	// connstr := "postgres", "user=postgres password="+os.Getenv("password")+" dbname=FlintCRUD sslmode=disable"
	connstr := os.Getenv("NEONTECH_URL")

	pq, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatalf("failed opening connection to posgresql: %v", err)
	}
	defer pq.Close()

	decoder := schema.NewDecoder()
	server.SchemaRegisterConverterNulls(decoder)

	query := db.New(pq)

	app := fiber.New(fiber.Config{})

	app.Use(logger.New())
	app.Get("/get", server.GetTrickHandler(ctx, query))
	app.Get("/list", server.ListTrickhandler(ctx, query))
	app.Post("/create", server.CreateTrickHandler(ctx, query))
	app.Delete("/delete", server.DeleteTrickHandler(ctx, query))

	log.Fatal(app.Listen(":4000"))
}

/*
	for {
		method := strings.ToLower(getAnswer("get, list, update, delete, or insert"))
		if method == "get" {
			id := getID()
			trick, _ := query.GetTrick(ctx, id)
			fmt.Println(trick)
			continue
		}
		if method == "list" {
			tricks, _ := query.GetAllTricks(ctx)
			fmt.Println(tricks)
			continue
		}
		if method == "update" {
			id := getID()

			name := getAnswer("name of trick")
			styleConv := getAnswer("style of trick")
			style, _ := strconv.Atoi(styleConv)
			powerConv := getAnswer("power of trick")
			var power bool
			if powerConv == "true" {
				power = true
			} else {
				power = false
			}

			_ = query.UpdateTrick(ctx, db.UpdateTrickParams{
				ID:    id,
				Name:  sql.NullString{String: name, Valid: true},
				Style: sql.NullInt32{Int32: int32(style), Valid: true},
				Power: sql.NullBool{Bool: power, Valid: true},
			})
			continue
		}
		if method == "delete" {
			id := getID()
			_ = query.DeleteTrick(ctx, id)
			continue
		}
		if method == "insert" {
			name := getAnswer("name of trick")
			styleConv := getAnswer("style of trick")
			style, _ := strconv.Atoi(styleConv)
			powerConv := getAnswer("power of trick")
			var power bool
			if powerConv == "true" {
				power = true
			} else {
				power = false
			}

			trick, _ := query.CreateTrick(ctx, db.CreateTrickParams{
				Name:  sql.NullString{String: name, Valid: true},
				Style: sql.NullInt32{Int32: int32(style), Valid: true},
				Power: sql.NullBool{Bool: power, Valid: true},
			})
			fmt.Println(trick)
			continue
		}
		os.Exit(0)
	}
}

func getAnswer(ask string) string {
	var input string
	fmt.Println(ask)
	fmt.Scanln(&input)
	return input
}

func getID() int32 {
	var input string
	fmt.Println("what id?")
	fmt.Scanln(&input)
	out, _ := strconv.Atoi(input)
	return int32(out)
}

*/
