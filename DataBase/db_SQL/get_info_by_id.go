package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func GetUserByID(ctx context.Context, conn pgx.Conn, id int) (User, error) {
	sqlQuery := `
	SELECT id, name, email, created_at
	FROM users
	WHERE ID = $1
	`

	var user User

	err := conn.QueryRow(ctx, sqlQuery, id).Scan(&user.ID, &user.Name, &user.Email, &user.created_at)
	if err != nil {
		return User{}, err
	}

	return user, err
}
