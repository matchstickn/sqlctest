package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	pq, err := sql.Open("postgres", "user=postgres password="+os.Getenv("password")+" dbname=FlintCRUD sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to posgresql: %v", err)
	}
	defer pq.Close()
}
