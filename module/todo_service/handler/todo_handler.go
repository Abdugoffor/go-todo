package todo_handler

import (
	"net/http"
	"strconv"
	todo_dto "task_app/module/todo_service/dto"
	todo_service "task_app/module/todo_service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TodoHandler struct {
	db          *gorm.DB
	todoService *todo_service.TodoService
}

func NewTodoHandler(group *echo.Group, db *gorm.DB) *TodoHandler {
	handler := TodoHandler{
		db:          db,
		todoService: todo_service.NewTodoService(db),
	}

	todoGroup := group.Group("/todo")
	todoGroup.GET("/", handler.All)
	todoGroup.POST("/create", handler.Create)
	todoGroup.GET("/:id", handler.Show)
	todoGroup.PUT("/update/:id", handler.Update)
	todoGroup.DELETE("/delete/:id", handler.Delete)
	todoGroup.DELETE("/delete-all", handler.DeleteAll)

	return &handler
}

func (h *TodoHandler) All(ctx echo.Context) error {
	userId := ctx.Get("user_id").(uint)

	todos, err := h.todoService.All(userId)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) Create(ctx echo.Context) error {

	var req todo_dto.CreateTodo

	if err := ctx.Bind(&req); err != nil {
		return err
	}
	userID := ctx.Get("user_id").(uint)
	req.UserID = userID

	todo, err := h.todoService.Create(&req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) Update(ctx echo.Context) error {
	var req todo_dto.UpdateTodo

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	idParam := ctx.Param("id")

	idUint64, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "ID noto‘g‘ri formatda"})
	}

	id := uint(idUint64)

	todo, err := h.todoService.Update(id, &req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) Show(ctx echo.Context) error {

	idParam := ctx.Param("id")

	idUint64, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "ID noto‘g‘ri formatda"})
	}

	id := uint(idUint64)

	todo, err := h.todoService.Show(id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) Delete(ctx echo.Context) error {

	idParam := ctx.Param("id")

	idUint64, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "ID noto‘g‘ri formatda"})
	}

	id := uint(idUint64)

	err = h.todoService.Delete(id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "Todo o‘chirildi"})
}

func (h *TodoHandler) DeleteAll(ctx echo.Context) error {
	userId := ctx.Get("user_id").(uint)

	err := h.todoService.DeleteAll(userId)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "Todos o‘chirildi"})
}
