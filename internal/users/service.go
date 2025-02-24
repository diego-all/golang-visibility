package users

import "golang-visibility/internal/models"

// Definimos la interfaz para el servicio de usuarios
type UserService interface {
	GetUsers() ([]models.User, error)
}

// Implementación del servicio de usuarios
type userService struct {
	repo UserRepository
}

// Constructor del servicio
func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

// Implementación de GetUsers
func (s *userService) GetUsers() ([]models.User, error) {
	return s.repo.GetAll()
}
