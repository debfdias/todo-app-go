package repository

import (
	"todo-server/internal/models"

	"gorm.io/gorm"
)

type TodoRepository struct {
    db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
    return &TodoRepository{db: db}
}

func (repo *TodoRepository) FindAll() ([]models.Todo, error) {
    var todos []models.Todo
    result := repo.db.Order("created_at desc").Find(&todos)
    return todos, result.Error
}

func (repo *TodoRepository) FindByID(id string) (*models.Todo, error) {
    var todo models.Todo
    result := repo.db.First(&todo, "id = ?", id)
    return &todo, result.Error
}

func (repo *TodoRepository) Create(todo *models.Todo) (*models.Todo, error) {
	result := repo.db.Create(todo)

	if result.Error != nil {
		return nil, result.Error
	}

	var createdTodo models.Todo
	err := repo.db.First(&createdTodo, "id = ?", todo.ID).Error
	return &createdTodo, err
}

func (repo *TodoRepository) Update(todo *models.Todo) error {
	result := repo.db.Save(todo)
	return result.Error
}

func (repo *TodoRepository) Delete(id string) error {
	result := repo.db.Delete(&models.Todo{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}