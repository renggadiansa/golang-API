package requests

import "time"

type UserRequest struct {
	Name     string    `json:"name" form:"name" binding:"required"`
	Email    string    `json:"email" form:"email" binding:"required"`
	Address  string    `json:"address" form:"address" binding:"required"`
	BornDate time.Time `json:"born_date" form:"born_date" binding:"required"`
}
