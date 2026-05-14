package database

import "time"

type User struct {
	ID         int
	Name       string
	Email      string
	created_at time.Time
}
