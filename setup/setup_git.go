package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Llamamos a la función que configura Git desde el paquete setup
	RunGitSetup()
}

// RunGitSetup configura git en tu máquina.
func RunGitSetup() {
	// Solicitar al usuario el alias, el email y el nombre de usuario de GitHub
	reader := bufio.NewReader(os.Stdin)

	// Solicitar el alias
	fmt.Print("Introduce tu alias de GitHub: ")
	alias, _ := reader.ReadString('\n')
	alias = strings.TrimSpace(alias) // Eliminar saltos de línea

	// Solicitar el email
	fmt.Print("Introduce tu email de GitHub: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	// Solicitar el username
	fmt.Print("Introduce tu username de GitHub: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	// Escribir "# projectName" en README.md
	file, err := os.OpenFile("README.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo README.md:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("# golang-visibility\n")
	if err != nil {
		fmt.Println("Error al escribir en README.md:", err)
		return
	}

	// Ejecutar git init
	cmd2 := exec.Command("git", "init")
	err = cmd2.Run()
	if err != nil {
		fmt.Println("Error al ejecutar git init:", err)
		return
	}

	// Ejecutar git branch -M main
	cmd3 := exec.Command("git", "branch", "-M", "main")
	err = cmd3.Run()
	if err != nil {
		fmt.Println("Error al ejecutar git branch -M main:", err)
		return
	}

	// Modificación en el comando de git remote add origin
	// Construir la URL usando el alias y username
	remoteURL := fmt.Sprintf("https://github.%s/%s/golang-visibility.git", alias, username)

	// Ejecutar git remote add origin con la nueva URL
	cmd4 := exec.Command("git", "remote", "add", "origin", remoteURL)
	err = cmd4.Run()
	if err != nil {
		fmt.Println("Error al ejecutar git remote add origin:", err)
		return
	}

	// Ejecutar git config --local user.email
	cmd5 := exec.Command("git", "config", "--local", "user.email", email)
	err = cmd5.Run()
	if err != nil {
		fmt.Println("Error al configurar email:", err)
		return
	}

	// Ejecutar git config --local user.name
	cmd6 := exec.Command("git", "config", "--local", "user.name", username)
	err = cmd6.Run()
	if err != nil {
		fmt.Println("Error al configurar user.name:", err)
		return
	}

	// Mensaje de éxito
	fmt.Println("Los comandos se han ejecutado correctamente.")
}
