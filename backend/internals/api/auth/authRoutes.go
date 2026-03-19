package auth

import "github.com/gin-gonic/gin"

func AuthRoutes(router *gin.RouterGroup, authController *AuthController) {
	auth := router.Group("/auth")
	auth.POST("/register", authController.RegisterUser)
	auth.POST("/login", authController.LoginUser)
	auth.POST("/reset-password", authController.ResetPassword)
	auth.POST("/otp-send")
	auth.POST("/otp-verify")
}