package utils

import (
	"log"
	"os"
	"os/exec"
)

// generateSwaggerDocs ejecuta swag init automáticamente en desarrollo
func GenerateSwaggerDocs() {
	cmd := exec.Command("swag", "init")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Println("Error generando Swagger docs:", err)
	}
}

// openSwaggerUI abre Swagger automáticamente en el navegador en desarrollo
func OpenSwaggerUI(osType, appPort string) {
	url := "http://localhost:" + appPort + "/swagger/index.html"

	var cmd *exec.Cmd
	switch osType {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url) // Windows
	case "darwin":
		cmd = exec.Command("open", url) // macOS
	case "linux":
		cmd = exec.Command("xdg-open", url) // Linux
	default:
		log.Println("Sistema operativo no compatible para abrir el navegador automáticamente.")
		return
	}

	err := cmd.Start()
	if err != nil {
		log.Println("Error abriendo Swagger UI:", err)
	}
}
