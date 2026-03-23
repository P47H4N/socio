package post

import (
	"github.com/P47H4N/socio/internals/middleware"
	"github.com/gin-gonic/gin"
)

func PostRoutes(router *gin.RouterGroup, postController *PostController) {
	postRoute := router.Group("/posts")
	{
		postRoute.GET("/:id", postController.GetPost)
		postRoute.GET("/user/:id", postController.GetUserPost)
		postRoute.GET("/:id/comments", postController.GetComment)
		postRoute.GET("/:id/reply", postController.GetReply)
		postPrivateRoute := postRoute.Group("/")
		postPrivateRoute.Use(middleware.UserMiddleware())
		{
			postPrivateRoute.GET("/", postController.Newsfeed)
			postPrivateRoute.POST("/", postController.CreatePost)
			postPrivateRoute.PATCH("/:id", postController.UpdatePost)
			postPrivateRoute.DELETE("/:id", postController.DeletePost)
			postPrivateRoute.POST("/:id/react", postController.ToggleReact)
			postPrivateRoute.POST("/:id/comments", postController.CreateComment)
			postPrivateRoute.DELETE("/comments/:id", postController.DeleteComment)
		}
	}
}
