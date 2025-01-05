CREATE INDEX IF NOT EXISTS idx_recommendation_user_id ON recommendation(user_id);

CREATE INDEX IF NOT EXISTS idx_recommendation_person_id ON recommendation(product_id);
