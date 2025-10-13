# backend/Dockerfile

# --- ETAPA 1: Builder ---
# Usa una imagen de Go para compilar el binario
FROM golang:1.25-alpine AS builder
WORKDIR /app

# Copia los archivos de dependencias de Go
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# --- ETAPA 2: Final ---
# Usa una imagen mínima para ejecutar el binario
FROM scratch
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]