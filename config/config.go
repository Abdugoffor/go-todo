package database

import (
	"fmt"
	"log"
	"os"
	auth_model "task_app/module/auth_service/model"
	todu_model "task_app/module/todo_service/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	user := os.Getenv("user")
	password := os.Getenv("password")
	host := os.Getenv("host")
	port := os.Getenv("port")
	dbname := os.Getenv("dbname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Bazaga ulanishda xatolik:", err)
	}

	DB = db
	err = DB.AutoMigrate(&auth_model.User{}, &todu_model.Todo{})
	if err != nil {
		log.Fatal("❌ AutoMigrate error:", err)
	}
	return db
}
