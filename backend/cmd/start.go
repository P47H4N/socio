package cmd

import (
	"github.com/P47H4N/socio/internals/api/auth"
	"github.com/P47H4N/socio/internals/api/setting"
	"github.com/P47H4N/socio/internals/api/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Start(router *gin.RouterGroup, db *gorm.DB) {
	// Auth
	authService := auth.NewService(db)
	authController := auth.NewController(authService)
	auth.AuthRoutes(router, authController)

	// User
	userService := user.NewService(db)
	userController := user.NewController(userService)
	user.UserRoutes(router, userController)


	// Report
	settingService := setting.NewService(db)
	settingController := setting.NewController(settingService)
	setting.SettingRoutes(router, settingController)
}