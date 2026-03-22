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
        return errors.New("Unable to delete user.")
    }
	return nil
}