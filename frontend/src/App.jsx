// frontend/src/App.jsx
import React, { useState, useEffect, useCallback } from 'react';
import axios from 'axios';
import './App.css';

function App() {
  const [url, setUrl] = useState('');
  const [scanId, setScanId] = useState(null);
  const [status, setStatus] = useState('');
  const [results, setResults] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');

  // Lógica de sondeo para obtener el estado y los resultados
  const pollScanStatus = useCallback(async () => {
    try {
      // Pide el estado actual del escaneo al backend
      const statusResponse = await axios.get(`/api/scans/${scanId}`);
      const currentStatus = statusResponse.data.status;
      setStatus(currentStatus);

      // Si el escaneo ha terminado, pide los resultados finales
      if (currentStatus === 'Completado') {
        setIsLoading(false);
        const resultsResponse = await axios.get(`/api/scans/${scanId}/results`);
        setResults(resultsResponse.data || []);
      }
    } catch (err) {
      console.error("Error durante el sondeo:", err);
      setError('Error al obtener el estado del escaneo.');
      setIsLoading(false);
    }
  }, [scanId]);

  // Hook que ejecuta el sondeo periódicamente
  useEffect(() => {
    if (isLoading && scanId && status !== 'Completado') {
      const intervalId = setInterval(pollScanStatus, 3000); // Sondea cada 3 segundos
      return () => clearInterval(intervalId); // Limpia el intervalo al salir
    }
  }, [isLoading, scanId, status, pollScanStatus]);


  // Manejador del envío del formulario
  const handleSubmit = async (event) => {
    event.preventDefault();
    setIsLoading(true);
    setResults([]);
    setStatus('');
    setError('');
    setScanId(null);

    try {
      // Envía la URL al backend para iniciar un nuevo escaneo
      const response = await axios.post('/api/scans', { url });
      setScanId(response.data.scan_id);
      setStatus('En cola'); // Actualiza el estado inicial
    } catch (err) {
      console.error("Error al iniciar el escaneo:", err);
      setError('No se pudo iniciar el escaneo. Revisa la URL o la conexión con la API.');
      setIsLoading(false);
    }
  };

  return (
    <div className="App">
      <header>
        <h1>Escáner de Inyección SQL </h1>
        <p>Full-Stack...</p>
      </header>
      <main>
        <form onSubmit={handleSubmit}>
          <input
            type="text"
            value={url}
            onChange={(e) => setUrl(e.target.value)}
            placeholder="Introduce la URL a escanear (ej: http://test.com/login.php?id=1)"
            required
          />
          <button type="submit" disabled={isLoading}>
            {isLoading ? 'Escaneando...' : 'Escanear'}
          </button>
        </form>

        {error && <p className="error-message">{error}</p>}

        {scanId && (
          <div className="status-container">
            <h3>Estado del Escaneo</h3>
            <p>ID: {scanId}</p>
            <p>Estado: <strong>{status}</strong></p>
          </div>
        )}

        {results.length > 0 && (
          <div className="results-container">
            <h3>Resultados </h3>
            {results.map((vuln, index) => (
              <div key={index} className="vulnerability-card">
                <p><strong>Tipo:</strong> {vuln.type}</p>
                <p><strong>URL:</strong> {vuln.url}</p>
                <p><strong>Payload Exitoso:</strong> <code>{vuln.payload}</code></p>
              </div>
            ))}
          </div>
        )}
         {status === 'Completado' && results.length === 0 && (
            <div className="results-container">
                <h3>Resultados</h3>
                <p>No se encontraron vulnerabilidades basadas en errores.</p>
            </div>
        )}
      </main>
    </div>
  );
}

export default App;