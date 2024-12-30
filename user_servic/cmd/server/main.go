package main

import (
	"log"
	"net"

	"user_service/internal/database"
	h "user_service/internal/handlers"
	"user_service/proto"

	"google.golang.org/grpc"
)

func main() {
	// Подключение к базе данных
	//	connStr := "postgres://isamutdin:dzhalal-05-@postgres:5432/recommendation_system?sslmode=disable"
	//	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", "isamutdin", "dzhalal-05-", "recommendation_system")

	db, err := database.NewConnection("host=postgres user=isamutdin password=a dbname=recommendation_system sslmode=disable")
	//db, err := database.NewConnection(connStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Настройка gRPC сервера
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	userServer := h.NewUserServer(db)
	proto.RegisterUserServiceServer(server, userServer)

	log.Println("Starting gRPC server on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
