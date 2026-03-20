package setting

import (
	"net/http"
	"strconv"

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

func (sc *SettingController) GetUserReports(c *gin.Context) {
	getUserId, _ := c.Get("userId")
	userId := getUserId.(uint)
	reports, err := sc.srv.GetUserReports(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Report fetch succussful.",
		Data: reports,
	})
}

func (sc *SettingController) GetReportDetails(c *gin.Context) {
	paramId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid report id.",
		})
		return
	}
	getUserId, _ := c.Get("userId")
	userId := getUserId.(uint)
	report, err := sc.srv.GetReportDetails(uint(paramId), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: false,
		Message: "Report fetch successfully.",
		Data: report,
	})
}

func (sc *SettingController) SubmitReport(c *gin.Context) {
	getUserId, _ := c.Get("userId")
	userId := getUserId.(uint)
	var reportBody ReportBody
	if err := c.ShouldBindBodyWithJSON(&reportBody); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid data.",
			Error: err.Error(),
		})
		return
	}
	if err := sc.srv.SubmitReport(userId, &reportBody); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Report submitted successfully.",
	})
}