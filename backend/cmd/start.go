package cmd

import (
	"github.com/P47H4N/socio/internals/api/auth"
	"github.com/P47H4N/socio/internals/api/message"
	"github.com/P47H4N/socio/internals/api/network"
	"github.com/P47H4N/socio/internals/api/post"
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

	// Post
	postService := post.NewService(db)
	postController := post.NewController(postService)
	post.PostRoutes(router, postController)

	// Network
	networkService := network.NewService(db)
	networkController := network.NewController(networkService)
	network.NetworkRoutes(router, networkController)

	// Message
	messageService := message.NewService(db)
	messageController := message.NewController(messageService)
	message.MessageRoutes(router, messageController)

	// Report
	settingService := setting.NewService(db)
	settingController := setting.NewController(settingService)
	setting.SettingRoutes(router, settingController)
}
