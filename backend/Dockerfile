# backend/Dockerfile

# Usamos una imagen oficial de Go como base. 'alpine' es una versión muy ligera.
FROM golang:1.25-alpine AS builder

# Establecemos el directorio de trabajo dentro del contenedor.
WORKDIR /app

# Copiamos los archivos de nuestro proyecto (que crearemos en la Fase 2) al contenedor.
COPY . .

# Este comando descargará las dependencias de Go.
RUN go mod tidy

# Compilamos nuestro código Go en un archivo ejecutable llamado 'main'.
RUN go build -o main .

# Este es el comando que se ejecutará cuando el contenedor arranque.
CMD ["./main"]