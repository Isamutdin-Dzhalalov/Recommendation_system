package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"user_service/proto"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type UserServer struct {
	proto.UnimplementedUserServiceServer
	db       *sql.DB
	producer *kafka.Producer
}

func NewUserServer(db *sql.DB, producer *kafka.Producer) *UserServer {
	return &UserServer{
		db:       db,
		producer: producer,
	}
}

func (s *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.UserResponse, error) {

	// TODO: добавить транзакцию, чтобы данные добавились и в бд и в kafka.
	query := `INSERT INTO users(name, email) VALUES($1, $2) RETURNING id`

	var id int32
	err := s.db.QueryRow(query, req.Name, req.Email).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	message := fmt.Sprintf(`{"id": %d, "name": "%s", "email": "%s"}`, id, req.Name, req.Email)

	topic := "user_update"
	err = s.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	if err != nil {
		log.Printf("failed to producer message: %v", err)
	}

	return &proto.UserResponse{
		Id:    id,
		Name:  req.Name,
		Email: req.Email,
	}, nil
}
