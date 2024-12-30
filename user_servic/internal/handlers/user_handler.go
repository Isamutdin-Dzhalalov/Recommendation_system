package handler

import (
	"context"
	"database/sql"
	"fmt"

	"user_service/proto"
)

type UserServer struct {
	proto.UnimplementedUserServiceServer
	db *sql.DB
}

func NewUserServer(db *sql.DB) *UserServer {
	return &UserServer{db: db}
}

func (s *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.UserResponse, error) {
	query := `INSERT INTO users(name, email) VALUES($1, $2) RETURNING id`
	var id int32
	err := s.db.QueryRow(query, req.Name, req.Email).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return &proto.UserResponse{
		Id:    id,
		Name:  req.Name,
		Email: req.Email,
	}, nil
}
