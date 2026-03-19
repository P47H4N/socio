package auth

type RegisterBody struct {
	Username string  `json:"username" binding:"required,min=3"`
	Email    string  `json:"email" binding:"required,email"`
	Phone    *string `json:"phone"`
	FullName string  `json:"full_name" binding:"required"`
	Password string  `json:"password" binding:"required,min=8,max=32"`
}

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResetBody struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	Password string `json:"password" binding:"required,min=8,max=32"`
}
