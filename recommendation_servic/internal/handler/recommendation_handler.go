package handler

import (
	"encoding/json"
	"log"
	"recommendation_service/internal/database"
	"recommendation_service/internal/messaging"
	"recommendation_service/internal/models"

	"github.com/nats-io/nats.go"
)

type RecommendationServer struct {
	db         *database.DB
	natsClient *messaging.NatsClient
}

func NewRecommendationServer(db *database.DB, natsClient *messaging.NatsClient) *RecommendationServer {
	return &RecommendationServer{db: db, natsClient: natsClient}
}

func (r *RecommendationServer) StartListening() {
	sub, err := r.natsClient.Conn.Subscribe("user_update", r.handleUserUpdate)
	if err != nil {
		log.Fatalf("failed to subscribe to user_update: %v", err)
	}

	defer sub.Unsubscribe()

	sub, err = r.natsClient.Conn.Subscribe("product_update", r.handleProductUpdate)
	if err != nil {
		log.Fatalf("failed to subscribe to product_update: %v", err)
	}

	defer sub.Unsubscribe()

	select {}
}

func (r *RecommendationServer) handleUserUpdate(msg *nats.Msg) {
	var user map[string]interface{}
	if err := json.Unmarshal(msg.Data, &user); err != nil {
		log.Printf("failed to unmarshal user update: %v", err)
		return
	}

	recommendations := r.generateRecommendations(user)

	for _, recommendation := range recommendations {
		_, err := r.db.Exec("INSERT INTO recommendations(user_id, product_id, score) VALUES($1, $2, $3)",
			recommendation.UserID, recommendation.ProductID, recommendation.Score)
		if err != nil {
			log.Printf("failed to insert recommendation: %v", err)
		}

	}
}

func (r *RecommendationServer) handleProductUpdate(msg *nats.Msg) {
	// TODO: logic update product
}

func (r *RecommendationServer) generateRecommendations(user map[string]interface{}) []models.Recommendation {
	var recommendations []models.Recommendation

	recommendations = append(recommendations, models.Recommendation{
		UserID:    int32(user["id"].(float64)),
		ProductID: 123,
		Score:     0.9,
	})
	return recommendations
}
