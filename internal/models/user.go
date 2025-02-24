package models

import "errors"

// Definimos la estructura del usuario
type User struct {
	ID    int
	Name  string
	Email string
}

// Simulamos una base de datos con un slice en memoria
type DB struct {
	users []User
}

// Función para inicializar la base de datos con datos de prueba
func NewDB() *DB {
	return &DB{
		users: []User{
			{ID: 1, Name: "Juan", Email: "juan@example.com"},
			{ID: 2, Name: "Maria", Email: "maria@example.com"},
		},
	}
}

// Método para obtener todos los usuarios
func (db *DB) GetAllUsers() ([]User, error) {
	if len(db.users) == 0 {
		return nil, errors.New("no hay usuarios")
	}
	return db.users, nil
}
