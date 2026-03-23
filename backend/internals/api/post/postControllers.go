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
			Message: "Invalid post id.",
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

func (pc *PostController) GetUserPost(c *gin.Context) {
	var userId uint = 0
	if val, exists := c.Get("userId"); exists {
		if id, ok := val.(uint); ok {
			userId = id
		}
	}
	paramId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid post id.",
		})
		return
	}
	posts, err := pc.srv.GetUserPost(uint(paramId), userId)
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
		Data: posts,
	})
}

func (pc *PostController) CreatePost(c *gin.Context) {
	var post PostBody
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid data.",
		})
		return
	}
	if post.Content == "" {
		_, fileErr := c.FormFile("media")
		if fileErr != nil {
			c.JSON(http.StatusBadRequest, models.Response{
				Success: false,
				Message: "Post content or media is required.",
			})
			return
		}
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

func (pc *PostController) UpdatePost(c *gin.Context) {
	var post PostBody
	paramId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid post id.",
		})
		return
	}
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid data.",
		})
		return
	}
	var mediaPath string
	file, err := c.FormFile("media")
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
	if err := pc.srv.UpdatePost(userId, uint(paramId), &post, mediaPath); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Post updated successfully.",
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
		Message: "Post deleted successfully.",
	})
}

func (pc *PostController) ToggleReact(c *gin.Context) {
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
	var react ReactBody
	if err := c.ShouldBindBodyWithJSON(&react); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid data.",
		})
		return
	}
	if err := pc.srv.ToggleReact(userId, uint(paramId), &react); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "React toggle.",
	})
}

func (pc *PostController) GetComment(c *gin.Context) {
	paramId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid post id.",
		})
		return
	}
	comments, err := pc.srv.GetComment(uint(paramId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Comments fetch successfully.",
		Data: comments,
	})
}

func (pc *PostController) GetReply(c *gin.Context) {
	paramId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid comment id.",
		})
		return
	}
	replies, err := pc.srv.GetReply(uint(paramId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Replies fetch successfully.",
		Data: replies,
	})
}

func (pc *PostController) CreateComment(c *gin.Context) {
	var comment CommentBody
	if err := c.ShouldBind(&comment); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid data.",
		})
		return
	}
	if comment.Content == "" {
		_, fileErr := c.FormFile("media")
		if fileErr != nil {
			c.JSON(http.StatusBadRequest, models.Response{
				Success: false,
				Message: "Comment content or media is required.",
			})
			return
		}
	}
	var mediaPath string
	file, err := c.FormFile("media")
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
	if err := pc.srv.CreateComment(userId, &comment, mediaPath); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Comment created successfully",
	})
}

func (pc *PostController) DeleteComment(c *gin.Context) {
	getUserId, _ := c.Get("userId")
	userId := getUserId.(uint)
	paramId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid comment id.",
		})
		return
	}
	if err := pc.srv.DeleteComment(userId, uint(paramId)); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Comment deleted successfully.",
	})
}