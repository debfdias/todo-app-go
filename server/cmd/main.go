package main

import (
	"todo-server/internal/handler"
	"todo-server/internal/repository"
	"todo-server/internal/service"
	"todo-server/pkg/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)
	healthHandler := handler.NewHealthHandler(db)

	api := router.Group("/api/")
	{
		api.GET("/health", healthHandler.CheckHealth)
		api.GET("/todos", todoHandler.GetTodos)
		api.POST("/todos", todoHandler.CreateTodo)
		api.GET("/todos/:id", todoHandler.GetTodo)
		api.PUT("/todos/:id", todoHandler.UpdateTodo)
		api.DELETE("/todos/:id", todoHandler.DeleteTodo)
	}

	return router
}

func main() {
    db := database.NewPostgresConnection()
    
    route := setupRoutes(db)
		route.Run(":8080")
}