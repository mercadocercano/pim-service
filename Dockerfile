# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Instalar dependencias de compilación
RUN apk add --no-cache git

# Copiar archivos de dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Final stage
FROM alpine:3.18

WORKDIR /app

# Instalar PostgreSQL client
RUN apk add --no-cache postgresql-client

# Crear directorios necesarios
RUN mkdir -p /app/scripts /app/migrations

# Copiar el binario compilado
COPY --from=builder /app/main .

# Copiar las migraciones
COPY --from=builder /app/src/category/infrastructure/persistence/migrations/*.sql /app/migrations/

# Exponer el puerto
EXPOSE 8080

# Ejecutar la aplicación directamente
CMD ["./main"] 