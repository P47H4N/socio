package message

import (
	"github.com/P47H4N/socio/internals/middleware"
	"github.com/gin-gonic/gin"
)

func MessageRoutes(router *gin.RouterGroup, messageController *MessageController) {
	messageRoute := router.Group("/messages")
	messageRoute.Use(middleware.UserMiddleware())
	{
		messageRoute.GET("/")           // Chat List
		messageRoute.GET("/:id")        // Chat History
		messageRoute.POST("/")          // Sent Message
		messageRoute.PATCH("/:id/read") // Seen Message
		messageRoute.DELETE("/:id")     // Delete Message
	}
}


// Upcoming