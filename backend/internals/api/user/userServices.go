package user

import (
	"errors"
	"fmt"
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

func (us *UserService) GetProfile(uid uint) (*models.User, error) {
	var user models.User
	query := us.db.First(&user, uid).Error
	if user.AccountStatus != "active" {
		return nil, helpers.AccountStatusCalculator(user.AccountStatus, user.DeletedAt.Time)
    }
	if query != nil {
		if query == gorm.ErrRecordNotFound {
			return nil, errors.New("User id not found.")
		}
		return nil, errors.New("Internal error.")
	}
	return &user, nil
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