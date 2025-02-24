package users

import "golang-visibility/internal/models"

// Interface para definir el contrato del repositorio
type UserRepository interface {
	GetAll() ([]models.User, error)
}

// Implementación del repositorio
type userRepository struct {
	db *models.DB
}

// Constructor del repositorio
func NewUserRepository(db *models.DB) UserRepository {
	return &userRepository{db: db}
}

// Implementación del método GetAll
func (r *userRepository) GetAll() ([]models.User, error) {
	return r.db.GetAllUsers()
}
