package main

import (
	"context"
	"log"
	"time"

	"user_service/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewUserServiceClient(conn)

	// Отправляем запрос на создание пользователя
	req := &proto.CreateUserRequest{
		Name:  "John Cache",
		Email: "john.cache@example.com",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.CreateUser(ctx, req)
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}

	log.Printf("Created User: ID=%d, Name=%s, Email=%s", res.Id, res.Name, res.Email)
}
