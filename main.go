package main

import (
	"bufio"
	"context"
	"os"
	database "users-cli/DataBase/db_SQL"
	"users-cli/chat"

	"github.com/jackc/pgx/v5"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:1234@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	if err := database.CreateTable(ctx, *conn); err != nil {
		panic(err)
	}

	chat.StartChat(ctx, *conn, scanner)
}
