package todo_cmd

import (
	auth_middleware "task_app/module/auth_service/middleware"
	todo_handler "task_app/module/todo_service/handler"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Cmd(router *echo.Echo, db *gorm.DB) {

	routerGroup := router.Group("/api", auth_middleware.JWTMiddleware)
	{
		todo_handler.NewTodoHandler(routerGroup, db)
	}
}
