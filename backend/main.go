package main

import (
	"fmt"
	"items/domain/items"
	"items/service/db/db"
	"items/service/server"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	pool, err := db.New()
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer pool.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5002))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	items.RegisterItemsServiceServer(s, server.New(pool))
	reflection.Register(s) // Register reflection service on gRPC server
	log.Println("Listening on port 5002...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
