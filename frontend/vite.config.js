import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  server: {
    host: true,          // Permite que Docker acceda al servidor dev
    port: 3000,          // Puerto que usarás en modo desarrollo
    proxy: {
      '/api': {
        target: 'http://backend:8080',
        changeOrigin: true,
        secure: false,
      }
    }
  },
  build: {
    outDir: 'dist',       // Donde se generará la build para Nginx
  }
})
