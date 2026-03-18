package auth

import (
	"errors"
	"log"

	"github.com/P47H4N/socio/internals/helpers"
	"github.com/P47H4N/socio/internals/models"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *AuthService{
	return &AuthService{
		db: db,
	}
}

func (as *AuthService) RegisterUser(body *RegisterBody) error {
	hashed_password, err := helpers.HashPassword(body.Password)
	if err != nil{
		log.Fatalln("Password Hash is not working. Password:", body.Password, "->", err)
	}
	if body.Phone != nil {
		if err := as.db.Where("phone = ?", body.Phone).First(&models.User{}).Error; err != nil{
			return errors.New("Mobile number already registered.")
		}
	}
	if body.Email != nil {
		if err := as.db.Where("email = ?", body.Email).First(&models.User{}).Error; err != nil{
			return errors.New("Email already registered.")
		}
	}
	newUser := &models.User{
		Username: body.Username,
		Email: *body.Email,
		Phone: *body.Phone,
		FullName: body.FullName,
		Password: hashed_password,
	}
	if err := as.db.Create(&newUser).Error; err != nil {
		return errors.New("Registration failed.")
	}
	return nil
}
