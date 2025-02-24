package handlers

import (
	"encoding/json"
	"golang-visibility/internal/users"
	"log"
	"net/http"
)

// Definimos el manejador
type UserHandler struct {
	service  users.UserService
	services users.UserRepository
}

// Constructor del manejador
func NewUserHandler(service users.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Método para manejar la petición GET /users
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers()
	if err != nil {
		log.Println(err)
		http.Error(w, "Error obteniendo usuarios", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
