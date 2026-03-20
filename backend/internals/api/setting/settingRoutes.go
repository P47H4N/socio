package setting

import (
	"github.com/P47H4N/socio/internals/middleware"
	"github.com/gin-gonic/gin"
)

func SettingRoutes(router *gin.RouterGroup, settingController *SettingController) {
	settingRoute := router.Group("/")
	settingRoute.Use(middleware.UserMiddleware())
	{
		settingRoute.GET("/settings", settingController.GetSetting)
		settingRoute.PATCH("/settings", settingController.UpdateSetting)
		settingRoute.GET("/reports", settingController.GetUserReports)
		settingRoute.GET("/reports/:id", settingController.GetReportDetails)
		settingRoute.POST("/reports", settingController.SubmitReport)
	}
}