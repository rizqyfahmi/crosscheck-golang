package param

type RegistrationParam struct {
	Name            string `form:"name" json:"name" validate:"required"`
	Email           string `form:"email" json:"email" validate:"required,email"`
	Password        string `form:"password" json:"password" validate:"required,eqfield=ConfirmPassword"`
	ConfirmPassword string `form:"confirmPassword" json:"confirmPassword" validate:"required,eqfield=Password"`
}
