package auth_cmd

import (
	auth_handler "task_app/module/auth_service/handler"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Cmd(router *echo.Echo, db *gorm.DB) {

	routerGroup := router.Group("/api")
	{
		auth_handler.NewAuthHandler(routerGroup, db)
	}
}
