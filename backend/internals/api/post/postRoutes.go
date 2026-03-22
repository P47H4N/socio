package post

import (
	"github.com/P47H4N/socio/internals/middleware"
	"github.com/gin-gonic/gin"
)

func PostRoutes(router *gin.RouterGroup, postController *PostController) {
	postRoute := router.Group("/posts")
	{
		postRoute.GET("/:id", postController.GetPost)
		postRoute.GET("/:id/comments") // View Comment
		postPrivateRoute := postRoute.Group("/")
		postPrivateRoute.Use(middleware.UserMiddleware())
		{
			postPrivateRoute.GET("/", postController.Newsfeed)
			postPrivateRoute.POST("/", postController.CreatePost)
			postPrivateRoute.PATCH("/:id") // Update Post
			postPrivateRoute.DELETE("/:id", postController.DeletePost)
			postPrivateRoute.POST("/:id/react")    // Toogle React
			postPrivateRoute.POST("/:id/comments") // Create Comment
			router.DELETE("/comments/:id")  // Delete Comment
		}
	}
}
