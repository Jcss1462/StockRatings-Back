package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jcss1462/StockRatings-Back/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Construcción de la conexión usando variables de entorno
func getDSN() (dsnDb string, dsnDefault string, dbName string) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
	dsnDb = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)
	dsnDefault = fmt.Sprintf("postgresql://%s:%s@%s:%s/defaultdb?sslmode=disable", user, password, host, port)
	// Si no existe la BD, conectar a `defaultdb`, de lo contrario, conectar a la BD definida
	return dsnDb, dsnDefault, dbName
}

func createDatabaseIfNotExists(dsnDefault, dbName string) {

	db, err := sql.Open("pgx", dsnDefault)
	if err != nil {
		log.Fatal("Error al conectar a CockroachDB:", err)
		fmt.Println("prueba")
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		log.Fatal("Error al crear la base de datos:", err)
	}

	fmt.Println("Base de datos creada o ya existente")
}

func ConnectDatabase(dsn string) {

	// Conectar a la base de datos
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Error al conectar a la base de datos:", err)
	}

	fmt.Println("✅ Conectado a CockroachDB")
	DB = db
}

func InitDB() {
	dsnDb, dsnDefault, dbName := getDSN()
	createDatabaseIfNotExists(dsnDefault, dbName)
	ConnectDatabase(dsnDb)

	// Crear tablas automáticamente
	DB.AutoMigrate(&models.Stock{})
}
