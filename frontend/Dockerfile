# frontend/Dockerfile

# --- ETAPA 1: Construcción (Builder) ---
# Usamos una imagen de Node.js para construir nuestro proyecto de React[cite: 143].
FROM node:18-alpine AS builder
WORKDIR /app

# Copiamos los archivos de dependencias y las instalamos.
COPY package*.json ./
RUN npm install

# Copiamos el resto del código fuente de React.
COPY . .

# Este comando crea la versión optimizada para producción de nuestra app[cite: 144].
RUN npm run build

# --- ETAPA 2: Producción (Final) ---
# Usamos una imagen de un servidor web súper ligero (Nginx) para servir los archivos[cite: 145].
FROM nginx:1.25-alpine

# Copiamos ÚNICAMENTE los archivos construidos de la etapa anterior.
# Esto hace que nuestra imagen final sea pequeña y segura[cite: 146].
COPY --from=builder /app/build /usr/share/nginx/html

# Borramos la configuración por defecto de Nginx.
RUN rm /etc/nginx/conf.d/default.conf

# Copiamos nuestro propio archivo de configuración de Nginx (lo crearemos a continuación).
COPY nginx.conf /etc/nginx/conf.d