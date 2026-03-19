package middleware

import (
	"net/http"

	"github.com/P47H4N/socio/internals/models"
	"github.com/P47H4N/socio/internals/helpers"
	"github.com/gin-gonic/gin"
)

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := helpers.GetToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.Response{
				Success: false,
				Message: "Unauthorized.",
			})
			c.Abort()
			return
		}
		claims, err := helpers.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.Response{
				Success: false,
				Message: "Unauthorized.",
			})
			c.Abort()
			return
		}
		c.Set("userId", claims.Id)
		c.Set("userName", claims.Username)
		c.Set("fullName", claims.FullName)
		c.Next()
	}
}