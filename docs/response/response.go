package docs

import "crosscheck-golang/app/features/authentication/domain/entity"

// The response when a request successfully processed
// swagger:response AuthSuccessResponse
type AuthSuccessResponse struct {
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
		Data entity.AuthEntity `json:"data"`
	}
}

// The response when a request fails to be processed caught by error content-type
// swagger:response InternalServerError
type InternalServerError struct {
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
// swagger:response BadRequest
type BadRequest struct {
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
