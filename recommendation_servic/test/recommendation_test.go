package test

import (
	"log"
	"testing"
	"time"

	"recommendation_service/internal/database"
	"recommendation_service/internal/handler"
	"recommendation_service/internal/messaging"
)

func TestRecommendationServer(t *testing.T) {

	connStr := "host=postgres user=user password=a dbname=recommendation_service sslmode=disable"

	db, err := database.NewConnection(connStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	natsURL := "nats://nats:4222"
	natsClient, err := messaging.NewNatsClient(natsURL)
	if err != nil {
		log.Fatalf("Failed to connect to nats: %v", err)
	}
	defer natsClient.Close()

	server := handler.NewRecommendationServer(db, natsClient)

	go server.StartListening()

	testUserUpdate := `{"id": 1, "name": "Test User"}`

	err = natsClient.Publisher("user_update", []byte(testUserUpdate))
	if err != nil {
		log.Fatalf("failed to publisher user update: %v", err)
	}

	time.Sleep(2 * time.Second)

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM recommendation_server WHERE userID = $1", 1).Scan(&count)
	if err != nil {
		log.Fatalf("failed to query database: %v", err)
	}

	if count == 0 {
		t.Errorf("No recommendation were generate  fo user")
	}

}
