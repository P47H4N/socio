package network

import (
	"net/http"
	"strconv"

	"github.com/P47H4N/socio/internals/models"
	"github.com/gin-gonic/gin"
)

type NetwrokController struct {
	srv *NetwrokService
}

func NewController(srv *NetwrokService) *NetwrokController {
	return &NetwrokController{
		srv: srv,
	}
}

func (nc *NetwrokController) GetFollowers(c *gin.Context) {
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
	followers, err := nc.srv.GetFollowers(userId, uint(paramId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Followers fetch successfully.",
		Data: followers,
	})
}

func (nc *NetwrokController) GetFollowing(c *gin.Context) {
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
	following, err := nc.srv.GetFollowing(userId, uint(paramId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Followings fetch successfully.",
		Data: following,
	})
}

func (nc *NetwrokController) Follow(c *gin.Context) {
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
	if err := nc.srv.Follow(userId, uint(paramId)); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User followed successfully.",
	})
}

func (nc *NetwrokController) Unfollow(c *gin.Context) {
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
	if err := nc.srv.Unfollow(userId, uint(paramId)); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User unfollowed successfully.",
	})
}

func (nc *NetwrokController) Block(c *gin.Context) {
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
	if err := nc.srv.Block(userId, uint(paramId)); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User blocked successfully.",
	})
}

func (nc *NetwrokController) Unblock(c *gin.Context) {
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
	if err := nc.srv.Unblock(userId, uint(paramId)); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User unblocked successfully.",
	})
}