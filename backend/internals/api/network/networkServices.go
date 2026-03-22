package network

import (
	"errors"

	"github.com/P47H4N/socio/internals/models"
	"gorm.io/gorm"
)

type NetwrokService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *NetwrokService {
	return &NetwrokService{
		db: db,
	}
}

func (ns *NetwrokService) GetFollowers(uid, id uint) ([]models.Follower, error) {
	var user models.UserSetting
	if err := ns.db.Where("user_id = ?", id).First(&user).Error; err != nil {
		return nil, errors.New("User not found.")
	}
	if user.IsPrivateAccount && uid != id {
		return nil, errors.New("This account is private.")
	}
	var followers []models.Follower
	if err := ns.db.Preload("Follower").Where("following_id = ?", id).Find(&followers).Error; err != nil {
		return nil, errors.New("Unable to fetch followers.")
	}
	return followers, nil
}

func (ns *NetwrokService) GetFollowing(uid, id uint) ([]models.Follower, error) {
	var user models.UserSetting
	if err := ns.db.Where("user_id = ?", id).First(&user).Error; err != nil {
		return nil, errors.New("User not found.")
	}
	if user.IsPrivateAccount && uid != id {
		return nil, errors.New("This account is private.")
	}
	var following []models.Follower
	if err := ns.db.Preload("Following").Where("follower_id = ?", id).Find(&following).Error; err != nil {
		return nil, errors.New("Unable to fetch followings.")
	}
	return following, nil
}

func (ns *NetwrokService) Follow(uid, id uint) error {
	if uid == id {
		return errors.New("You can't follow yourself.")
	}
	var follower models.Follower
	if err := ns.db.Where("follower_id = ? AND following_id = ?", uid, id).First(&follower).Error; err == nil {
		return errors.New("Already following this user.")
	}
	newFollow := models.Follower{
		FollowingID: id,
		FollowerID:  uid,
	}
	if err := ns.db.Create(&newFollow).Error; err != nil {
		return errors.New("Unable to follow.")
	}
	return nil
}

func (ns *NetwrokService) Unfollow(uid, id uint) error {
	var follower models.Follower
	if err := ns.db.Where("follower_id = ? AND following_id = ?", uid, id).First(&follower).Error; err != nil {
		return errors.New("You should follow first to unfollow.")
	}
	if err := ns.db.Unscoped().Delete(&follower).Error; err != nil {
		return errors.New("Unable to unfollow at this moment.")
	}
	return nil
}

func (ns *NetwrokService) Block(uid, id uint) error {
	if uid == id {
		return errors.New("You can't block yourself.")
	}

	// No longer wait for block.

	var block models.BlockList
	if err := ns.db.Where("blocker_id = ? AND blocked_id = ?", uid, id).First(&block).Error; err == nil {
		return errors.New("Alreary block this user.")
	}
	newBlock := models.BlockList{
		BlockerID: uid,
		BlockedID: id,
	}
	if err := ns.db.Create(&newBlock).Error; err != nil {
		return errors.New("Unable to block this user.")
	}
	return nil
}

func (ns *NetwrokService) Unblock(uid, id uint) error {
	var block models.BlockList
	if err := ns.db.Where("blocker_id = ? AND blocked_id = ?", uid, id).First(&block).Error; err != nil {
		return errors.New("You should block first to unblock.")
	}
	if err := ns.db.Unscoped().Delete(&block).Error; err != nil {
		return errors.New("Unable to unblock at this moment.")
	}
	return nil
}
