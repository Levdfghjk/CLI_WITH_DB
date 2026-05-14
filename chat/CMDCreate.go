package chat

import (
	"bufio"
	"errors"
	"fmt"
)

func CMDCreate(scanner *bufio.Scanner) (string, string, error) {
	fmt.Println("Введитt ваше имя:")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Введитt ваше email:")
	scanner.Scan()
	email := scanner.Text()

	if email == "" || name == "" {
		return "", "", errors.New("Имя или Email пустые!")
	}

	return name, email, nil
}
