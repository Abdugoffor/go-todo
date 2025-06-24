package todu_model

import (
	auth_model "task_app/module/auth_service/model"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	UserID uint
	User   auth_model.User `gorm:"foreignKey:UserID"`
	Date   time.Time       `json:"date"`
	Text   string          `json:"text"`
	Status bool            `json:"status" gorm:"default:true"`
}

func (Todo) TableName() string {
	return "todos"
}
