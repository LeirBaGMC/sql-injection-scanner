// backend/scanner/scanner.go
package scanner

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/LeirBaGMC/sql-scanner/database" // Importa nuestro paquete de base de datos
)

// RunScan ejecuta la lógica de detección de vulnerabilidades
func RunScan(scanID string, targetURL string) {
	log.Printf("Iniciando escaneo %s para la URL: %s", scanID, targetURL)
	database.DB.Exec("UPDATE scans SET status = ? WHERE id = ?", "En ejecución", scanID)

	errorPayloads := []string{"'", "\"", ";", " OR 1=1 --"}
	errorSignatures := []string{"You have an error in your SQL syntax", "unclosed quotation mark", "mysql_fetch"}

	client := &http.Client{Timeout: 10 * time.Second}
	vulnerabilityFound := false

	for _, payload := range errorPayloads {
		req, err := http.NewRequest("GET", targetURL+payload, nil)
		if err != nil {
			continue
		}

		resp, err := client.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		bodyBytes, _ := io.ReadAll(resp.Body)
		bodyString := string(bodyBytes)

		for _, signature := range errorSignatures {
			if strings.Contains(bodyString, signature) {
				log.Printf("¡VULNERABILIDAD ENCONTRADA! Tipo: Basado en Errores, Payload: %s", payload)
				database.DB.Exec("INSERT INTO vulnerabilities (scan_id, url, type, payload) VALUES (?, ?, ?, ?)",
					scanID, targetURL, "Basado en Errores", payload)
				vulnerabilityFound = true
				break
			}
		}
		if vulnerabilityFound {
			break
		}
	}

	database.DB.Exec("UPDATE scans SET status = ? WHERE id = ?", "Completado", scanID)
	log.Printf("Escaneo %s completado.", scanID)
}
