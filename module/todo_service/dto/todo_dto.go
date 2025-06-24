package todo_dto

import "time"

type CreateTodo struct {
	Date   time.Time `json:"date" validate:"required"`
	Text   string    `json:"text" validate:"required"`
	Status bool      `json:"status"`
	UserID uint      `json:"user_id" validate:"required"`
}

type UpdateTodo struct {
	Date   time.Time `json:"date" validate:"required"`
	Text   string    `json:"text" validate:"required"`
	Status bool      `json:"status"`
	UserID uint      `json:"user_id" validate:"required"`
}

type TodoResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Date      time.Time `json:"date"`
	Text      string    `json:"text"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
