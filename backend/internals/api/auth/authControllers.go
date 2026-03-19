package auth

import (
	"net/http"

	"github.com/P47H4N/socio/internals/helpers"
	"github.com/P47H4N/socio/internals/models"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	srv *AuthService
}

func NewController(srv *AuthService) *AuthController {
	return &AuthController{
		srv: srv,
	}
}

func (ac *AuthController) RegisterUser(c *gin.Context) {
	var registerBody RegisterBody
	if err := c.ShouldBindBodyWithJSON(&registerBody); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid data.",
			Error:   err.Error(),
		})
		return
	}
	if len(registerBody.Password) < 8 || len(registerBody.Password) > 32 {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Password must be between 8 and 32 characters.",
		})
		return
	}
	if registerBody.Phone != nil {
		if len(*registerBody.Phone) < 10 {
			c.JSON(http.StatusBadRequest, models.Response{
				Success: false,
				Message: "Invalid mobile number.",
			})
			return
		}
	}
	if _, err := helpers.ValidateUsername(registerBody.Username); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Error: err.Error(),
		})
		return
	}
	if err := ac.srv.RegisterUser(&registerBody); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Registration successful.",
	})
}

func (ac *AuthController) LoginUser(c *gin.Context) {
	var loginBody LoginBody
	if err := c.ShouldBindBodyWithJSON(&loginBody); err != nil{
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid data.",
			Error: err.Error(),
		})
		return
	}
	token, user, err := ac.srv.LoginUser(&loginBody)
	if err != nil{
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Login successful.",
		Data: gin.H{
			"token": token,
			"user": user,
		},
	})
}

func (ac *AuthController) ResetPassword(c *gin.Context) {
	var email struct{
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindBodyWithJSON(&email); err != nil{
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Please provide a valid email address.",
			Error: err.Error(),
		})
		return
	}
	token := helpers.GenerateVerificationToken(8)
	if err := ac.srv.ResetPassword(email.Email, token); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Password reset successful.",
	})
}