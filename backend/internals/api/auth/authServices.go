package auth

import (
	"errors"
	"log"
	"time"

	"github.com/P47H4N/socio/internals/helpers"
	"github.com/P47H4N/socio/internals/models"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (as *AuthService) RegisterUser(body *RegisterBody) error {
	hashed_password, err := helpers.HashPassword(body.Password)
	if err != nil {
		log.Fatalln("Password Hash is not working. Password:", body.Password, "->", err)
	}
	if body.Phone != nil {
		if err := as.db.Where("phone = ?", body.Phone).First(&models.User{}).Error; err != nil {
			return errors.New("Mobile number already registered.")
		}
	}
	if body.Email != nil {
		if err := as.db.Where("email = ?", body.Email).First(&models.User{}).Error; err != nil {
			return errors.New("Email already registered.")
		}
	}
	newUser := &models.User{
		Username: body.Username,
		Email:    *body.Email,
		Phone:    *body.Phone,
		FullName: body.FullName,
		Password: hashed_password,
	}
	if err := as.db.Create(&newUser).Error; err != nil {
		return errors.New("Registration failed.")
	}
	return nil
}

func (as *AuthService) LoginUser(body *LoginBody) (string, *models.User, error) {
	var user models.User
	if err := as.db.Where("username = ?", body.Username).Or("email = ?", body.Username).Or("mobile = ?", body.Username).First(&user).Error; err != nil {
		return "", nil, errors.New("User not found.")
	}
	if !helpers.CheckPasswordHash(body.Password, user.Password) {
		return "", nil, errors.New("Wrong password.")
	}
	token, err := helpers.GenerateToken(&models.TokenBody{
		Id:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})
	if err != nil {
		return "", nil, err
	}
	return token, &user, nil
}
