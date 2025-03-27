# Usar imagen base de Go
FROM golang:1.21-alpine AS builder

# Establecer directorio de trabajo
WORKDIR /app

# Copiar archivos de go mod y sum
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar codigo fuente
COPY . .

# Construir la aplicacion
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Imagen final mas ligera
FROM alpine:latest

WORKDIR /root/

# Copiar binario compilado
COPY --from=builder /app/main .

# Exponer puerto
EXPOSE 8080

# Variables de entorno por defecto
ENV DB_HOST=localhost \
    DB_PORT=5432 \
    DB_USER=postgres \
    DB_PASSWORD=password \
    DB_NAME=laliga

# Comando de inicio
CMD ["./main"]