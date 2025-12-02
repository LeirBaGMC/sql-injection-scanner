// backend/main.go
package main

import (
	"log"

	"github.com/LeirBaGMC/sql-scanner/api" 
	"github.com/LeirBaGMC/sql-scanner/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Conectar a la base de datos
	database.Connect()
	defer database.DB.Close()

	// 2. Crear las tablas
	database.CreateTables()

	// 3. Configurar el router de la API
	router := gin.Default()
	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/scans", api.StartScanHandler)
		apiGroup.GET("/scans/:id", api.GetScanStatusHandler)
		apiGroup.GET("/scans/:id/results", api.GetScanResultsHandler)
	}

	// 4. Iniciar el servidor
	log.Println("Iniciando el servidor en el puerto 8080...")
	router.Run(":8080")
}
