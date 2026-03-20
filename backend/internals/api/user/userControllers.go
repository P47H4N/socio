package user

import (
	"net/http"
	"strconv"

	"github.com/P47H4N/socio/internals/models"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	srv *UserService
}

func NewController(srv *UserService) *UserController {
	return &UserController{
		srv: srv,
	}
}

func (uc *UserController) GetProfile(c *gin.Context) {
	paramId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid user id.",
		})
		return
	}
	getUserId, _ := c.Get("userId")
	userId := getUserId.(uint)
	if uint(paramId) != userId {
		c.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Unauthorized.",
		})
		return
	}
	user, err := uc.srv.GetProfile(uint(paramId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User found.",
		Data:    user,
	})
}

func (uc *UserController) ChangePassword(c *gin.Context) {
	getUserId, _ := c.Get("userId")
	userId := getUserId.(uint)
	var passwordBody ChangePasswordBody
	if err := c.ShouldBindBodyWithJSON(&passwordBody); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid Data.",
		})
		return
	}
	if err := uc.srv.ChangePassword(&passwordBody, userId); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: false,
		Message: "Password Changed Successfully.",
	})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	getUserId, _ := c.Get("userId")
	userId := getUserId.(uint)
	paramId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid user id.",
		})
		return
	}
	if uint(paramId) != userId {
		c.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Unauthorized.",
		})
		return
	}
	if err := uc.srv.DeleteUser(uint(paramId)); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User deleted successfully.",
	})
}
