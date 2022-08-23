package entity

// swagger:model AuthEntity
type AuthEntity struct {
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTUxNjIzOTAyMn0.bM1Td-Z3cEH17gygGwbXUefCN7NaKEXazy3khKzwjj0
	// required: true
	AccessToken string `json:"access_token"`
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTUxNjIzOTAyMn0.bM1Td-Z3cEH17gygGwbXUefCN7NaKEXazy3khKzwjj0
	// required: true
	RefreshToken string `json:"refresh_token"`
}
