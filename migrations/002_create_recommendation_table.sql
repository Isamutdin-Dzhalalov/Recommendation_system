CREATE TABLE IF NOT EXISTS recommendation (
	id SERIAL PRIMARY KEY,
	user_id INTEGER REFERANCEC users(id),
	product_id INTEGER NOT NULL,
	score DECIMAL NOT NULL,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW()
);
