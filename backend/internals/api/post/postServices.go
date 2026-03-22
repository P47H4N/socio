package post

import (
	"errors"
	"time"

	"github.com/P47H4N/socio/internals/models"
	"gorm.io/gorm"
)

type PostService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *PostService {
	return &PostService{
		db: db,
	}
}

func (ps *PostService) Newsfeed() ([]*models.Post, error) {
	return nil, nil
}

func (ps *PostService) GetPost(id uint) (*models.Post, error) {
	var post models.Post
	if err := ps.db.Preload("User").Where("id = ?", id).First(&post).Error; err != nil {
		return nil, errors.New("Post not found.")
	}
	return &post, nil
}

func (ps *PostService) CreatePost(uid uint, body *PostBody, path string) error {
	post := models.Post{
		UserID: uid,
		Content: body.Content,
		MediaURL: path,
		PostType: body.PostType,
		Privacy: body.Privacy,
	}
	if err := ps.db.Create(&post).Error; err != nil {
		return errors.New("Unable to create post.")
	}
	return nil
}

func (ps *PostService) UpdatePost(uid, pid uint, body *PostBody, path string) error {
	var post models.Post
	if err := ps.db.Where("id = ? AND user_id = ?", pid, uid).First(&post).Error; err != nil {
		return errors.New("Post not found.")
	}
	updateData := make(map[string]interface{})
	if body.Content != "" {
        updateData["content"] = body.Content
    }
    if body.PostType != "" {
        updateData["post_type"] = body.PostType
    }
    if body.Privacy != "" {
        updateData["privacy"] = body.Privacy
    }
    if path != "" {
        updateData["media_url"] = path
    }
	if len(updateData) == 0 {
        return errors.New("Nothing is updated.") 
    }
	if err := ps.db.Model(&post).Updates(updateData).Error; err != nil {
		return errors.New("Upable to update post data.")
	}
	return nil
}

func (ps *PostService) DeletePost(uid, pid uint) error {
	var post models.Post
	if err := ps.db.First(&post, pid).Error; err != nil {
		return errors.New("Post not found.")
	}
	if post.UserID != uid {
		return errors.New("You are not authorized to delete this post.")
	}
	post.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}
	if err := ps.db.Save(&post).Error; err != nil {
        return errors.New("Unable to delete post.")
    }
	return nil
}

func (ps *PostService) ToggleReact(uid, pid uint, body *ReactBody) error {
	var react models.React
	if err := ps.db.Where("user_id = ? AND post_id = ?", uid, pid).First(&react).Error; err == nil {
		if body.Type == "" || body.Type == react.Type {
			if err := ps.db.Delete(&react).Error; err != nil {
				return errors.New("Unable to toggle reaction.")
			}
			return nil
		}
		react.Type = body.Type
		if err := ps.db.Save(&react).Error; err != nil {
			return errors.New("Unable to change reaction.")
		}
		return nil
	}
	newReact := models.React{
		UserID: uid,
		PostID: pid,
		Type: body.Type,
	}
	if err := ps.db.Create(&newReact).Error; err != nil {
		return errors.New("Unable to create reaction.")
	}
	return nil
}

func (ps *PostService) GetComment(pid uint) ([]models.Comment, error) {
	var comments []models.Comment
	if err := ps.db.Preload("User").Where("post_id = ? AND parent_id IS NULL", pid).Order("created_at DESC").Find(&comments).Error; err != nil {
		return nil, errors.New("Unable to find comments.")
	}
	return comments, nil
}

func (ps *PostService) GetReply(cid uint) ([]models.Comment, error) {
	var replies []models.Comment
	if err := ps.db.Preload("User").Where("parent_id = ?", cid).Order("created_at DESC").Find(&replies).Error; err != nil {
		return nil, errors.New("Unable to find replies.")
	}
	return nil, nil
}

func (ps *PostService) CreateComment(uid uint, body *CommentBody, path string) error {
	comment := models.Comment{
		PostID: body.PostID,
		UserID: uid,
		ParentID: body.ParentID,
		Content: body.Content,
		MediaURL: path,
	}
	if err := ps.db.Create(&comment).Error; err != nil {
        return errors.New("Unable to create comment.")
    }
	return nil
}

func (ps *PostService) DeleteComment(uid, cid uint) error {
	var comment models.Comment
	if err := ps.db.First(&comment, cid).Error; err != nil {
		return errors.New("Comment not found.")
	}
	if comment.UserID != uid {
		return errors.New("You are not authorized to delete this comment.")
	}
	comment.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}
	if err := ps.db.Save(&comment).Error; err != nil {
        return errors.New("Unable to delete comment.")
    }
	return nil
}