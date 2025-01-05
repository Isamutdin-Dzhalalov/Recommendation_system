package main

import (
	"log"
	"net"
	"recommendation_service/internal/database"
	"recommendation_service/internal/handler"
	"recommendation_service/internal/messaging"

	"google.golang.org/grpc"
)

func main() {
	connStr := "host=postgres user=isamutdin password=a db=recommendation_ses sslmode=disable"

	db, err := database.NewConnection(connStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	natsClient, err := messaging.NewNatsClient("nats://nats:4222")
	if err != nil {
		log.Fatalf("failed to connect to nats: %v", err)
	}

	defer natsClient.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	recommendationServer := handler.NewRecommendationServer(db, natsClient)
	go recommendationServer.StartListening()

	log.Println("Starting gRPC server on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
