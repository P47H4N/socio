package auth

type RegisterBody struct {
	Username string  `json:"username" binding:"required,min=3"`
	Email    *string `json:"email" binding:"email"`
	Phone    *string `json:"phone"`
	FullName string  `json:"full_name" binding:"required"`
	Password string  `json:"-" binding:"required,min=6"`
}
