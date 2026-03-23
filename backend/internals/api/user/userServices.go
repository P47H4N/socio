package user

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/P47H4N/socio/internals/helpers"
	"github.com/P47H4N/socio/internals/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (us *UserService) GetProfile(identifier string) (*models.User, int64, int64, error) {
	var user models.User
	var follower, following int64
	var query *gorm.DB

	id, err := strconv.Atoi(identifier)
	if err == nil {
		query = us.db.First(&user, id)
	} else {
		query = us.db.Where("username = ?", identifier).First(&user)
	}

	if query.Error != nil {
		if query.Error == gorm.ErrRecordNotFound {
			return nil, 0, 0, errors.New("User not found.")
		}
		return nil, 0, 0, errors.New("Internal error.")
	}

	if user.AccountStatus != "active" {
		return nil, 0, 0, helpers.AccountStatusCalculator(user.AccountStatus, user.DeletedAt.Time)
	}

	us.db.Model(&models.Follower{}).Where("following_id = ?", user.ID).Count(&follower)
	us.db.Model(&models.Follower{}).Where("follower_id = ?", user.ID).Count(&following)

	return &user, follower, following, nil
}

func (us *UserService) ChangePassword(body *ChangePasswordBody, uid uint) error {
	user := &models.User{}
	if err := us.db.Model(&user).Where("id = ?", uid).Error; err != nil {
		return errors.New("User not found.")
	}
	if !helpers.CheckPasswordHash(body.OldPassword, user.Password) {
		return errors.New("Wrong old password.")
	}
	newHashedPassword, _ := helpers.HashPassword(body.NewPassword)
	if err := us.db.Model(&user).Update("password", newHashedPassword).Error; err != nil {
        return errors.New("Failed to update password.")
    }
	return nil
}

func (us *UserService) DeleteUser(uid uint) error {
	var user models.User
	if err := us.db.First(&user, uid).Error; err != nil {
		return errors.New("User not found.")
	}
	user.AccountStatus = "deleted";
	user.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}
	user.Email = fmt.Sprintf("%s_del_%d", user.Email, time.Now().Unix())
	if err := us.db.Save(&user).Error; err != nil {
        return errors.New("Unable to delete user.")
    }
	return nil
}