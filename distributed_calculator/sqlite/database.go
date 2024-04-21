package sqlite

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

// CreateDatabase создает базу данных SQLite.
func CreateDatabase() error {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URL is not set")
	}

	pool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		return fmt.Errorf("pgxpool.Connect: %v", err)
	}
	defer pool.Close()

	schema :=
		CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		login TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS expressions (
		id UUID PRIMARY KEY,
		user_id INTEGER NOT NULL REFERENCES users (id),
		expression TEXT NOT NULL,
		result FLOAT NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT
