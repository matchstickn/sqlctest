package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

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

	pq, err := sql.Open("postgres", "user=postgres password="+os.Getenv("password")+" dbname=FlintCRUD sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to posgresql: %v", err)
	}
	defer pq.Close()

	myHomeIsABasketballInHighRoad247 := db.New(pq)

	tricks, err := myHomeIsABasketballInHighRoad247.GetTrick(ctx, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tricks)
}
