# frontend/Dockerfile

# --- ETAPA 1: Construcción (Builder) ---
# Usa una imagen de Node.js para construir el proyecto
FROM node:20-alpine AS builder
WORKDIR /app

# Copia los archivos de dependencias de Node.js, no de Go
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# --- ETAPA 2: Producción (Final) ---
# Usa Nginx para servir los archivos estáticos
FROM nginx:1.25-alpine
# La carpeta de salida de Vite es 'dist'
COPY --from=builder /app/dist /usr/share/nginx/html
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d