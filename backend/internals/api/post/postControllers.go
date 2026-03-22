package post

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/P47H4N/socio/internals/models"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	srv *PostService
}

func NewController(srv *PostService) *PostController {
	return &PostController{
		srv: srv,
	}
}

func (pc *PostController) Newsfeed(c *gin.Context) {

}

func (pc *PostController) GetPost(c *gin.Context) {
	paramId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid user id.",
		})
		return
	}
	post, err := pc.srv.GetPost(uint(paramId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Post fetch successfully.",
		Data: post,
	})
}

func (pc *PostController) CreatePost(c *gin.Context) {
	var post PostBody
	if err := c.ShouldBindBodyWithJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid data.",
		})
		return
	}
	file, err := c.FormFile("media")
	var mediaPath string
	if err == nil {
        filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
        dst := "./uploads/" + filename
        if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{
				Success: false,
				Message: "Failed to save media.",
			})
            return
        }
        mediaPath = "/uploads/" + filename
    }
	getUserId, _ := c.Get("userId")
	userId := getUserId.(uint)
	if err := pc.srv.CreatePost(userId, &post, mediaPath); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Post created successfully.",
	})
}

func (pc *PostController) DeletePost(c *gin.Context) {
	getUserId, _ := c.Get("userId")
	userId := getUserId.(uint)
	paramId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid post id.",
		})
		return
	}
	if err := pc.srv.DeletePost(userId, uint(paramId)); err != nil {
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