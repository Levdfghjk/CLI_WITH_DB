package chat

import (
	"context"
	"fmt"
	database "users-cli/DataBase/db_SQL"

	"github.com/jackc/pgx/v5"
)

func StartChat(ctx context.Context, conn pgx.Conn) {
	for {
		var cmd string

		fmt.Println("Введите команду:")

		n, err := fmt.Scan(&cmd)
		if n == 0 && err != nil {
			fmt.Println("Ошибка чтения команды!")
			return
		}

		if cmd == "create" {
			name, email := CMDCreate()
			err2 := database.CreateUser(ctx, conn, name, email)
			if err2 != nil {
				panic(err2)
			}
		} else if cmd == "end" {
			break
		} else {
			fmt.Println("Неизвестная команда!")
		}
	}
}
