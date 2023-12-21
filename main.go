package main

import (
	"bill-splitting/api"
	"bill-splitting/auth"
	db "bill-splitting/db/sqlc"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	authSecret := os.Getenv("AUTH_SECRET")
	dbDriver := "postgres"
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")

	conn, err := sql.Open(dbDriver, fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", dbDriver, dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	auth.NewAuth()

	store := db.NewStore(conn)
	server := api.NewServer(store, authSecret)
	server.Start("0.0.0.0:8080")
}
