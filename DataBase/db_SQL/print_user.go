package database

import "fmt"

func PrintUser(user User) {
	fmt.Println("ID:", user.ID)
	fmt.Println("Имя:", user.Name)
	fmt.Println("Email:", user.Email)
	fmt.Println("Created At:", user.created_at)
}
