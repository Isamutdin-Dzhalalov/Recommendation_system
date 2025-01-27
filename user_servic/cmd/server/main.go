package main

import (
	"log"
	"net"

	"user_service/internal/database"
	h "user_service/internal/handlers"
	"user_service/internal/messaging"
	"user_service/proto"

	"google.golang.org/grpc"
)

func main() {

	db, err := database.NewConnection("host=postgres user=isamutdin password=a dbname=recommendation_system sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	natsClient, err := messaging.NewNatsClient("nats://nats:4222")
	if err != nil {
		log.Fatalf("failed to connect nats: %v", err)
	}
	defer natsClient.Close()

	// Настройка gRPC сервера
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	userServer := h.NewUserServer(db, natsClient)
	proto.RegisterUserServiceServer(server, userServer)

	log.Println("Starting gRPC server on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
