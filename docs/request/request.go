package docs

import "crosscheck-golang/app/features/authentication/data/param"

// swagger:parameters registration
type RegistrationParameter struct {
	// in: body
	// required: true
	Body param.RegistrationParam
}

// swagger:parameters login
type LoginParameter struct {
	// in: body
	// required: true
	Body param.LoginParam
}
