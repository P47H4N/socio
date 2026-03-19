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
	var phoneStr string
	if body.Phone != nil {
		phoneStr = *body.Phone
		if err := as.db.Where("phone = ?", phoneStr).First(&models.User{}).Error; err == nil {
			return errors.New("Mobile number already registered.")
		}
	}
	if err := as.db.Where("email = ?", body.Email).First(&models.User{}).Error; err == nil {
		return errors.New("Email already registered.")
	}
	newUser := &models.User{
		Username: body.Username,
		Email:    body.Email,
		Phone:    phoneStr,
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
	if err := as.db.Where("username = ?", body.Username).Or("email = ?", body.Username).Or("phone = ?", body.Username).First(&user).Error; err != nil {
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

func (as *AuthService) ForgotPassword(email, token string) error {
	expiration := time.Now().Add(15 * time.Minute)
	securityData := models.Security{
		Email:     email,
		Token:     token,
		Type:      "password_reset",
		ExpiredAt: expiration,
		IsUsed:    false,
	}
	as.db.Model(&models.Security{}).Where("email = ? AND type = ? AND is_used = ?", email, "password_reset", false).Update("is_used", true)
	if err := as.db.Create(&securityData).Error; err != nil {
		return errors.New("Reset token can not be stored.")
	}
	return nil
}

func (as *AuthService) ConfirmToken(email, token string) error {
	result := as.db.Model(&models.Security{}).Where("email = ? AND token = ? AND is_used = ?", email, token, false).Update("is_used", true)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Token invalid or already used")
	}
	return nil
}

func (as *AuthService) ResetPassword(reset *ResetBody) error {
	var security models.Security
	if err := as.db.Where("email = ? AND token = ?", reset.Email, reset.Token).First(&security).Error; err != nil {
		return errors.New("Invalid email or token.")
	}
	if time.Now().After(security.ExpiredAt) {
		return errors.New("Token has expired.")
	}
	hashedPassword, err := helpers.HashPassword(reset.Password)
	if err != nil {
		return errors.New("Failed to hash password.")
	}
	update := map[string]any{
		"password": hashedPassword,
	}
	if err := as.db.Model(&models.User{}).Where("email = ?", reset.Email).Updates(&update).Error; err != nil {
		return errors.New("Password update failed.")
	}
	return nil
}
