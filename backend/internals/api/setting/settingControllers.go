package setting

import (
	"net/http"

	"github.com/P47H4N/socio/internals/models"
	"github.com/gin-gonic/gin"
)

type SettingController struct {
	srv *SettingService
}

func NewController(srv *SettingService) *SettingController {
	return &SettingController{
		srv: srv,
	}
}

func (sc *SettingController) GetSetting(c *gin.Context) {
	getUserId, _ := c.Get("userId")
	userId := getUserId.(uint)
	userSetting, err := sc.srv.GetSetting(userId);
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User settings found.",
		Data: userSetting,
	})
}

func (sc *SettingController) UpdateSetting(c *gin.Context) {
	getUserId, _ := c.Get("userId")
	userId := getUserId.(uint)
	var settingsBody SettingsBody
	if err := c.ShouldBindBodyWithJSON(&settingsBody); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid data.",
			Error: err.Error(),
		})
		return
	}
	if err := sc.srv.UpdateSetting(&settingsBody, userId); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Successfully updated data.",
	})
}