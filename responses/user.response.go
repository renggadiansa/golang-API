package responses

type UserResponse struct {
	ID *int `json:"id"`
	Name *string `json:"name"`
	Address *string `json:"address"`
}