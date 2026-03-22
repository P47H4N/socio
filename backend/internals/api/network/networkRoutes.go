package network

import (
	"github.com/P47H4N/socio/internals/middleware"
	"github.com/gin-gonic/gin"
)

func NetworkRoutes(router *gin.RouterGroup, networkController *NetwrokController) {
	networkRoute := router.Group("/")
	networkRoute.Use(middleware.UserMiddleware())
	{
		networkRoute.GET("/:id/followers", networkController.GetFollowers)
		networkRoute.GET("/:id/following", networkController.GetFollowing)
		networkRoute.POST("/follow/:id", networkController.Follow)
		networkRoute.DELETE("/unfollow/:id", networkController.Unfollow)
		networkRoute.POST("/block/:id", networkController.Block)
		networkRoute.POST("/unblock/:id", networkController.Unblock)
	}
}