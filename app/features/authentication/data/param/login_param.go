package param

// swagger:model LoginParameter
type LoginParam struct {
	// The username of a user
	// example: johndoe@mail.com
	// required: true
	Username string `form:"username" json:"username" validate:"required"`
	// The password of a user
	// example: HelloPassword
	// required: true
	Password string `form:"password" json:"password" validate:"required"`
}
