// backend/database/database.go
package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// DB es la conexión a la base de datos que será accesible por otros paquetes
var DB *sql.DB

// Connect se encarga de inicializar la conexión
func Connect() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", user, password, host, dbname)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error fatal al configurar la conexión a la DB: %v", err)
	}

	for i := 0; i < 10; i++ {
		err = DB.Ping()
		if err == nil {
			log.Println("Conexión a la base de datos exitosa.")
			return
		}
		log.Println("Esperando a la base de datos...")
		time.Sleep(3 * time.Second)
	}
	log.Fatalf("No se pudo conectar a la base de datos después de varios intentos: %v", err)
}

// CreateTables crea las tablas si no existen
func CreateTables() {
	scansTable := `
    CREATE TABLE IF NOT EXISTS scans (
        id VARCHAR(36) PRIMARY KEY, url TEXT NOT NULL, status VARCHAR(20) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`
	vulnerabilitiesTable := `
    CREATE TABLE IF NOT EXISTS vulnerabilities (
        id INT AUTO_INCREMENT PRIMARY KEY, scan_id VARCHAR(36), url TEXT,
        type VARCHAR(50), payload TEXT, FOREIGN KEY (scan_id) REFERENCES scans(id)
    );`

	if _, err := DB.Exec(scansTable); err != nil {
		log.Fatalf("Error al crear la tabla 'scans': %v", err)
	}
	if _, err := DB.Exec(vulnerabilitiesTable); err != nil {
		log.Fatalf("Error al crear la tabla 'vulnerabilities': %v", err)
	}
	log.Println("Tablas de la base de datos verificadas/creadas exitosamente.")
}
