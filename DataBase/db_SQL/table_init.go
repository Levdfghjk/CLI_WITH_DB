package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
	);
	`

	_, err := conn.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}

	return err
}
