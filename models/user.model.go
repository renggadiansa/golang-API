package models

import "time"

type User struct {
	ID       *int       `json:"id"`
	Name     *string    `json:"name"`
	Email    *string    `json:"email"`
	Address  *string    `json:"address"`
	BornDate *time.Time `json:"born_date"`
}
