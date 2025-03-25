package handler

import (
	"net/http"
	"todo-server/internal/models"
	"todo-server/internal/service"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
    service *service.TodoService
}

func NewTodoHandler(service *service.TodoService) *TodoHandler {
    return &TodoHandler{service: service}
}

func (h *TodoHandler) GetTodos(c *gin.Context) {
    todos, err := h.service.GetAllTodos()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) GetTodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := h.service.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if todo == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTodo, err := h.service.CreateTodo(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTodo)
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.service.UpdateTodo(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteTodo(id)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo deleted successfully"})
}
