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

// The response when a request successfully processed
// swagger:response RegistrationSuccessResponse
type _ struct {
	// in: body
	// required: true
	Body struct {
		// example: success
		// required: true
		Status string `json:"status"`
		// example: Request successfully processed
		// required: true
		Message string `json:"message"`
		// required: true
		Data AuthEntity `json:"data"`
	}
}

// The response when a request fails to be processed caught by error content-type
// swagger:response Registration500Response
type _ struct {
	// in: body
	// required: true
	Body struct {
		// example: error
		// required: true
		Status string `json:"status"`
		// example: Internal server error
		// required: true
		Message string `json:"message"`
		// example: null
		// required: true
		Data interface{} `json:"data"`
	}
}

// The response when a request fails to be processed caught by error validation, encryption, database, access token, and refresh token
// swagger:response Registration400Response
type _ struct {
	// in: body
	// required: true
	Body struct {
		// example: error
		// required: true
		Status string `json:"status"`
		// example: Bad request
		// required: true
		Message string `json:"message"`
		// example: null
		// required: true
		Data interface{} `json:"data"`
	}
}
