package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func GetTenLastUsers(ctx context.Context, conn pgx.Conn) error {
	sqlQuery := `
	SELECT id, name, email, created_at
	FROM users
	ORDER BY id DESC
	LIMIT 10;
	`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, email string
		var created_at time.Time

		if err := rows.Scan(&id, &name, &email, &created_at); err != nil {
			return err
		}

		fmt.Println("ID:", id)
		fmt.Println("name:", name)
		fmt.Println("email:", email)
		fmt.Println("created at:", created_at)
		fmt.Println("---------------------------------------")
	}

	return nil
}
