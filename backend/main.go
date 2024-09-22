package main

import (
	"bill-splitting/api"
	"bill-splitting/auth"
	"bill-splitting/db"
	"bill-splitting/gapi"
	"bill-splitting/proto"
	"bill-splitting/token"
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

	auth.NewAuth()

	store := db.NewStore()
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
