package main

import (
	"bill-splitting/api"
	"bill-splitting/auth"
	db "bill-splitting/db/sqlc"
	"bill-splitting/gapi"
	"bill-splitting/proto"
	"bill-splitting/token"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	tokenMaker := token.NewJWTMaker(authSecret)
	server := api.NewServer(store, tokenMaker)
	go server.Start("0.0.0.0:8080")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	grpcServer := gapi.NewServer(store, tokenMaker)
	proto.RegisterBillSplittingServer(s, grpcServer)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
