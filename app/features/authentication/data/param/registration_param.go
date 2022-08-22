package param

// swagger:model RegistrationParameter
type RegistrationParam struct {
	// The full name of a new user
	// example: John Doe
	// required: true
	Name string `form:"name" json:"name" validate:"required"`
	// The email address of a new user
	// example: johndoe@mail.com
	// required: true
	Email string `form:"email" json:"email" validate:"required,email" example:"Some name"`
	// The password of a new user
	// example: HelloPassword
	// required: true
	Password string `form:"password" json:"password" validate:"required,eqfield=ConfirmPassword"`
	// the confirmation password of a new user
	// example: HelloPassword
	// required: true
	ConfirmPassword string `form:"confirmPassword" json:"confirmPassword" validate:"required,eqfield=Password"`
}

// swagger:parameters registration
type _ struct {
	// in: body
	// required: true
	Body RegistrationParam
}
