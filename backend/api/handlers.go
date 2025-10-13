// backend/api/handlers.go
package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/LeirBaGMC/sql-scanner/database" // Importamos los paquetes locales
	"github.com/LeirBaGMC/sql-scanner/scanner"
	"github.com/gin-gonic/gin"
)

func StartScanHandler(c *gin.Context) {
	var json struct {
		URL string `json:"url" binding:"required"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El campo 'url' es requerido"})
		return
	}

	scanID := fmt.Sprintf("%d", time.Now().UnixNano())

	_, err := database.DB.Exec("INSERT INTO scans (id, url, status) VALUES (?, ?, ?)", scanID, json.URL, "En cola")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar el escaneo en la base de datos"})
		return
	}

	go scanner.RunScan(scanID, json.URL)

	c.JSON(http.StatusAccepted, gin.H{"scan_id": scanID})
}

func GetScanStatusHandler(c *gin.Context) {
	scanID := c.Param("id")
	var status string
	err := database.DB.QueryRow("SELECT status FROM scans WHERE id = ?", scanID).Scan(&status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Escaneo no encontrado"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": status})
}

func GetScanResultsHandler(c *gin.Context) {
	scanID := c.Param("id")
	rows, err := database.DB.Query("SELECT url, type, payload FROM vulnerabilities WHERE scan_id = ?", scanID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los resultados de la base de datos"})
		return
	}
	defer rows.Close()

	var results []gin.H
	for rows.Next() {
		var url, vulnType, payload string
		if err := rows.Scan(&url, &vulnType, &payload); err != nil {
			log.Printf("Error al escanear fila de resultados: %v", err)
			continue
		}
		results = append(results, gin.H{"url": url, "type": vulnType, "payload": payload})
	}
	c.JSON(http.StatusOK, results)
}
