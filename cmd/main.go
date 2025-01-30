package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/matchstickn/sqlctest/assets/db"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// connstr := "postgres", "user=postgres password="+os.Getenv("password")+" dbname=FlintCRUD sslmode=disable"
	connstr := os.Getenv("NEONTECH_URL")

	pq, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatalf("failed opening connection to posgresql: %v", err)
	}
	defer pq.Close()

	query := db.New(pq)

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
