package main

import (
	"log"
	database "task_app/config"
	auth_cmd "task_app/module/auth_service"
	todo_cmd "task_app/module/todo_service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Xatolik: .env fayl topilmadi")
	}

	db := database.ConnectDB()

	router := echo.New()

	auth_cmd.Cmd(router, db)
	todo_cmd.Cmd(router, db)

	router.Logger.Fatal(router.Start(":8090"))
}

