package auth_service

import (
	"os"
	auth_dto "task_app/module/auth_service/dto"
	auth_model "task_app/module/auth_service/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{DB: db}
}

func (s *AuthService) Register(input auth_dto.RegisterInput) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := auth_model.User{Name: input.Name, Email: input.Email, Password: string(hashed)}
	return s.DB.Create(&user).Error
}

func (s *AuthService) Login(input auth_dto.LoginInput) (string, error) {
	var user auth_model.User
	if err := s.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return "", err
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
