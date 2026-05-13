package chat

import "fmt"

func CMDCreate() (string, string) {
	fmt.Println("Введитt ваше имя:")

	var name string
	var email string

	n, err := fmt.Scan(&name)
	if n == 0 && err != nil {
		fmt.Println("Ошибка чтения команды!")
		return "", ""
	}

	if name == "" {
		fmt.Println("Ошибка! Имя не должно быть пустым!")
	}

	fmt.Println("Введитt ваш email:")
	n, err1 := fmt.Scan(&email)
	if n == 0 && err1 != nil {
		fmt.Println("Ошибка чтения команды!")
		return "", ""
	}

	if email == "" {
		fmt.Println("Ошибка! Имя не должно быть пустым!")
	}

	return name, email
}
