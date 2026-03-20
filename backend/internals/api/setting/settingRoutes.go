package setting

import (
	"github.com/P47H4N/socio/internals/middleware"
	"github.com/gin-gonic/gin"
)

func SettingRoutes(router *gin.RouterGroup, settingController *SettingController) {
	settingRoute := router.Group("/")
	settingRoute.Use(middleware.UserMiddleware())
	{
		router.GET("/settings", settingController.GetSetting)
		router.PATCH("/settings", settingController.UpdateSetting)
		router.GET("/reports/:id")
		router.POST("/reports")
	}
}