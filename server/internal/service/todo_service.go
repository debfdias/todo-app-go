package service

import (
	"errors"
	"todo-server/internal/models"
	"todo-server/internal/repository"

	"github.com/google/uuid"
)

type TodoService struct {
    repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
    return &TodoService{repo: repo}
}

func (service *TodoService) GetAllTodos() ([]models.Todo, error) {
    return service.repo.FindAll()
}

func (service *TodoService) GetTodo(id string) (*models.Todo, error) {
    return service.repo.FindByID(id)
}

func (service *TodoService) CreateTodo(input *models.Todo) (*models.Todo, error) {
	if input.Title == "" {
		return nil, errors.New("title cannot be empty")
	}

	todo := models.Todo{
		ID:          uuid.New().String(),
		Title:       input.Title,
		Description: input.Description,
		Completed:   false,
	}
	createdTodo, err := service.repo.Create(&todo)

	if err != nil {
		return nil, err
	}
	
	return createdTodo, nil
}

func (service *TodoService) UpdateTodo(id string, input models.Todo) (*models.Todo, error) {
	existing, err := service.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("todo not found")
	}

	if input.Title != "" {
		existing.Title = input.Title
	}
	if input.Description != "" {
		existing.Description = input.Description
	}
	existing.Completed = input.Completed

	err = service.repo.Update(existing)
	return existing, err
}

func (service *TodoService) DeleteTodo(id string) error {
	return service.repo.Delete(id)
}