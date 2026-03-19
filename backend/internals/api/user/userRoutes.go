package user

import (
	"github.com/P47H4N/socio/internals/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, userController *UserController) {
	userRoute := router.Group("/users")
	userRoute.Use(middleware.UserMiddleware())
	{
		userRoute.GET("/:id", userController.GetProfile)
		userRoute.GET("/search", )
		userRoute.POST("/change-password", userController.ChangePassword)
		userRoute.PATCH("/update/bio")
		userRoute.PATCH("/update/name")
		userRoute.PATCH("/update/username")
		userRoute.POST("/update/avatar")
		userRoute.POST("/update/cover")
		userRoute.DELETE("/:id", userController.DeleteUser)
	}
}
