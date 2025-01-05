package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"user_service/internal/messaging"
	"user_service/proto"
)

type UserServer struct {
	proto.UnimplementedUserServiceServer
	db         *sql.DB
	natsClient *messaging.NatsClient
}

func NewUserServer(db *sql.DB, natsClient *messaging.NatsClient) *UserServer {
	return &UserServer{db: db, natsClient: natsClient}
}

func (s *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.UserResponse, error) {
	query := `INSERT INTO users(name, email) VALUES($1, $2) RETURNING id`
	var id int32
	err := s.db.QueryRow(query, req.Name, req.Email).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	message := map[string]interface{}{
		"id":    id,
		"name":  req.Name,
		"email": req.Email,
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal message: %v", err)
	}

	if err := s.natsClient.Publisher("user_update", messageBytes); err != nil {
		return nil, fmt.Errorf("failed to publish message: %v", err)
	}

	return &proto.UserResponse{
		Id:    id,
		Name:  req.Name,
		Email: req.Email,
	}, nil
}
