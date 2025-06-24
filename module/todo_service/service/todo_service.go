package todo_service

import (
	todo_dto "task_app/module/todo_service/dto"
	todu_model "task_app/module/todo_service/model"

	"gorm.io/gorm"
)

type TodoService struct {
	db *gorm.DB
}

func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{db: db}
}

func (s *TodoService) GetByID(id uint) (*todo_dto.TodoResponse, error) {
	var todoModel todu_model.Todo
	if err := s.db.First(&todoModel, id).Error; err != nil {
		return nil, err
	}

	return &todo_dto.TodoResponse{
		ID:        todoModel.ID,
		UserID:    todoModel.UserID,
		Date:      todoModel.Date,
		Text:      todoModel.Text,
		Status:    todoModel.Status,
		CreatedAt: todoModel.CreatedAt,
		UpdatedAt: todoModel.UpdatedAt,
	}, nil
}

func (s *TodoService) All(userID uint) ([]todo_dto.TodoResponse, error) {
	var responses []todo_dto.TodoResponse
	err := s.db.Model(&todu_model.Todo{}).
		Select("id, user_id, date, text, status, created_at, updated_at").
		Where("user_id = ?", userID).
		Scan(&responses).Error

	return responses, err
}

func (s *TodoService) Show(id uint) (*todo_dto.TodoResponse, error) {
	var todoModel todu_model.Todo
	if err := s.db.First(&todoModel, id).Error; err != nil {
		return nil, err
	}

	return &todo_dto.TodoResponse{
		ID:        todoModel.ID,
		UserID:    todoModel.UserID,
		Date:      todoModel.Date,
		Text:      todoModel.Text,
		Status:    todoModel.Status,
		CreatedAt: todoModel.CreatedAt,
		UpdatedAt: todoModel.UpdatedAt,
	}, nil
}

func (s *TodoService) Create(todoDto *todo_dto.CreateTodo) (*todo_dto.TodoResponse, error) {
	todoModel := todu_model.Todo{
		Date:   todoDto.Date,
		Text:   todoDto.Text,
		Status: todoDto.Status,
		UserID: todoDto.UserID,
	}

	if err := s.db.Create(&todoModel).Error; err != nil {
		return nil, err
	}

	return &todo_dto.TodoResponse{
		ID:        todoModel.ID,
		UserID:    todoModel.UserID,
		Date:      todoModel.Date,
		Text:      todoModel.Text,
		Status:    todoModel.Status,
		CreatedAt: todoModel.CreatedAt,
		UpdatedAt: todoModel.UpdatedAt,
	}, nil
}

func (s *TodoService) Update(id uint, todoDto *todo_dto.UpdateTodo) (*todo_dto.TodoResponse, error) {
	var todoModel todu_model.Todo
	if err := s.db.First(&todoModel, id).Error; err != nil {
		return nil, err
	}

	todoModel.Date = todoDto.Date
	todoModel.Text = todoDto.Text
	todoModel.Status = todoDto.Status

	if err := s.db.Save(&todoModel).Error; err != nil {
		return nil, err
	}

	return &todo_dto.TodoResponse{
		ID:        todoModel.ID,
		UserID:    todoModel.UserID,
		Date:      todoModel.Date,
		Text:      todoModel.Text,
		Status:    todoModel.Status,
		CreatedAt: todoModel.CreatedAt,
		UpdatedAt: todoModel.UpdatedAt,
	}, nil
}

func (s *TodoService) Delete(id uint) error {
	var todoModel todu_model.Todo
	return s.db.Delete(&todoModel, id).Error
}

func (s *TodoService) DeleteAll(userID uint) error {
	return s.db.Where("user_id = ?", userID).Delete(&todu_model.Todo{}).Error
}
