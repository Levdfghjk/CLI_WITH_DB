package chat

import (
	"errors"
	"fmt"
)

func CMDInfo() (int, error) {
	fmt.Println("Введите ID:")

	var ID int

	n, err := fmt.Scan(&ID)
	if n == 0 && err != nil {
		fmt.Println("Ошибка чтения команды!")
		return 0, errors.New("Ошибка чтения команды")
	}

	return ID, nil
}
