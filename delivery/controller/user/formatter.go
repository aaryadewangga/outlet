package user

// =========== REQUEST ============ //

type CreateUserRequestFormat struct {
	Name     string `json:"name" form:"name" validate:"required,min=3,max=20,excludesall=!@#?^#*()_+-=0123456789%&"`
	User_uid string
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=3,max=15"`
	// Image    string `json:"image" form:"image"`
}

// =========== RESPONSE ============ //

type UserCreateResponse struct {
	User_uid string `json:"user_uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Roles    bool   `json:"roles" form:"roles"`
	// Image    string `json:"image" form:"image"`
}
