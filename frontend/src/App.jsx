import React, { useState, useEffect, useCallback } from 'react';
import axios from 'axios';
import './App.css'; 

function App() {
  const [url, setUrl] = useState('');
  const [scanId, setScanId] = useState(null);
  const [status, setStatus] = useState('');
  const [results, setResults] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [apiError, setApiError] = useState(''); 
  const [formError, setFormError] = useState('');

  const handleUrlChange = (e) => {
    const newUrl = e.target.value;
    setUrl(newUrl);

    if (newUrl === '') {
        setFormError(''); 
        return;
    }


    const urlPattern = /^(https?:\/\/)?([\da-z\.-]+|localhost)(:\d+)?([\/\w \.-?=&]*)*\/?$/;

    if (!urlPattern.test(newUrl)) {
        setFormError('La URL no parece válida.');
    } else {
        setFormError(''); 
    }
  };


  const pollScanStatus = useCallback(async () => {
    try {
      const statusResponse = await axios.get(`/api/scans/${scanId}`);
      const currentStatus = statusResponse.data.status;
      setStatus(currentStatus);

      if (currentStatus === 'Completado') {
        setIsLoading(false);
        const resultsResponse = await axios.get(`/api/scans/${scanId}/results`);
        setResults(resultsResponse.data || []);
      }
    } catch (err) {
      console.error("Error durante el sondeo:", err);
      setApiError('Error al obtener el estado del escaneo.');
      setIsLoading(false);
    }
  }, [scanId]);

  useEffect(() => {
    if (isLoading && scanId && status !== 'Completado') {
      const intervalId = setInterval(pollScanStatus, 3000); 
      return () => clearInterval(intervalId);
    }
  }, [isLoading, scanId, status, pollScanStatus]);


  
  const handleSubmit = async (event) => {
    event.preventDefault(); 
    if (formError || url === '') {
      setFormError('Por favor, introduce una URL válida para escanear.');
      return; 
    }
    setIsLoading(true);
    setResults([]);
    setStatus('');
    setApiError('');
    setScanId(null);
    try {
      const response = await axios.post('/api/scans', { url });
      setScanId(response.data.scan_id);
      setStatus('En cola'); 
    } catch (err) {
      console.error("Error al iniciar el escaneo:", err);
      setApiError('No se pudo iniciar el escaneo. Revisa la URL o la conexión con la API.');
      setIsLoading(false);
    }
  };

 
  return (
    <div className="App">
      <header>
        <h1>Escáner de Inyección SQL </h1>
        <p>Full-Stack</p>
      </header>
      <main>
        <form onSubmit={handleSubmit}>
          <input
            type="url"
            value={url}
            onChange={handleUrlChange}
            placeholder="Introduce la URL a escanear (ej: http://localhost:8000)"
            required
            style={formError ? { borderColor: '#dc3545', outlineColor: '#dc3545' } : {}}
          />
          <button 
            type="submit" 
            disabled={isLoading || !!formError} // !!formError convierte el string en booleano
          >
            {isLoading ? 'Escaneando...' : 'Escanear'}
          </button>
        </form>
        
        {formError && <p className="error-message" style={{marginTop: '-20px', marginBottom: '20px'}}>{formError}</p>}
        {apiError && <p className="error-message">{apiError}</p>}
        
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
                <p>No se encontraron vulnerabilidades.</p>
            </div>
        )}
      </main>
    </div>
  );
}

export default App;