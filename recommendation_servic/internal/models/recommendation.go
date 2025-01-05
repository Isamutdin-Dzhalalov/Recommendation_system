package models

type Recommendation struct {
	UserID    int32   `json:"user_id"`
	ProductID int32   `json:"product_id"`
	Score     float64 `json:"score"`
}
