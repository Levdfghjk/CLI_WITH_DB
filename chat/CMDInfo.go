package chat

import (
	"bufio"
	"fmt"
	"strconv"
)

func CMDInfo(scanner *bufio.Scanner) (int, error) {
	fmt.Println("Введите ID:")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, err
	}
	return id, nil
}
