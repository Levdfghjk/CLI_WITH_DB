package chat

import (
	"bufio"
	"context"
	"fmt"
	database "users-cli/DataBase/db_SQL"

	"github.com/jackc/pgx/v5"
)

func StartChat(ctx context.Context, conn pgx.Conn, scanner *bufio.Scanner) {
	for {
		fmt.Println("")
		fmt.Println("Введите команду:")
		scanner.Scan()
		cmd := scanner.Text()

		if cmd == "create" {
			name, email, err := CMDCreate(scanner)
			if err != nil {
				panic(err)
			}

			err2 := database.CreateUser(ctx, conn, name, email)
			if err2 != nil {
				panic(err2)
			}
		} else if cmd == "end" {
			break
		} else if cmd == "get info" {
			id, err := CMDInfo()
			if err != nil {
				panic(err)
			}

			u, err := database.GetUserByID(ctx, conn, id)
			if err != nil {
				panic(err)
			}

			database.PrintUser(u)
		} else if cmd == "last 10" {
			if err := database.GetTenLastUsers(ctx, conn); err != nil {
				panic(err)
			}
		} else {
			fmt.Println("Неизвестная команда!")
		}
	}
}
