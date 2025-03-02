package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error cargando el archivo .env", err)
	}

}
