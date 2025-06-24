package auth_handler

import (
	"net/http"
	auth_dto "task_app/module/auth_service/dto"
	auth_service "task_app/module/auth_service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db          *gorm.DB
	authService *auth_service.AuthService
}

func NewAuthHandler(group *echo.Group, db *gorm.DB) {
	handler := AuthHandler{
		db:          db,
		authService: auth_service.NewAuthService(db),
	}

	authGroup := group.Group("/auth")
	authGroup.POST("/register", handler.Register)
	authGroup.POST("/login", handler.Login)
}

func (h *AuthHandler) Register(ctx echo.Context) error {
	var input auth_dto.RegisterInput
	
	if err := ctx.Bind(&input); err != nil {
		return err
	}
	
	err := h.authService.Register(input)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "Foydalanuvchi yaratildi"})
}

func (h *AuthHandler) Login(c echo.Context) error {
	var input auth_dto.LoginInput
	if err := c.Bind(&input); err != nil {
		return err
	}
	token, err := h.authService.Login(input)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Email yoki parol xato"})
	}
	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
