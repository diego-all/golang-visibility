package main

import (
	"fmt"
	"golang-visibility/internal/handlers"
	"golang-visibility/internal/models"
	"golang-visibility/internal/users"
	"log"
	"net/http"
)

func main() {
	// Inicializar la base de datos (simulada)
	db := models.NewDB()

	// Crear repositorio y servicio de usuarios
	userRepo := users.NewUserRepository(db)
	userService := users.NewUserService(userRepo)

	// Crear manejador de usuarios
	userHandler := handlers.NewUserHandler(userService)

	// Definir rutas
	mux := http.NewServeMux()
	mux.HandleFunc("/users", userHandler.GetAllUsers)

	fmt.Println("Servidor corriendo en :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
