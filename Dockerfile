# Multi-stage build para optimizar el tamaño de la imagen final
FROM golang:1.24-alpine AS builder

# Instalar certificados SSL
RUN apk --no-cache add ca-certificates

# Establecer directorio de trabajo
WORKDIR /app

# Copiar archivos de dependencias primero (para aprovechar cache de Docker)
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Verificar las dependencias
RUN go mod verify

# Copiar el código fuente
COPY . .

# Compilar la aplicación con optimizaciones específicas para Go 1.23
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -a -installsuffix cgo \
    -ldflags="-w -s" \
    -o noc-monitoring \
    ./main.go

# Segunda etapa: imagen final mínima
FROM alpine:latest

# Instalar herramientas mínimas necesarias
RUN apk --no-cache add ca-certificates tzdata

# Crear usuario no-root para seguridad
RUN adduser -D -s /bin/sh soapuser

# Crear directorios para logs, datos y configuración
RUN mkdir -p /app/logs /app/data /app/config && \
    chown -R soapuser:soapuser /app

# Establecer directorio de trabajo
WORKDIR /app

# Copiar el binario compilado desde la etapa anterior
COPY --from=builder /app/noc-monitoring .

# Copiar certificados SSL
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Cambiar propiedad del binario al usuario no-root
RUN chown soapuser:soapuser noc-monitoring && \
    chmod +x noc-monitoring

# Cambiar al usuario no-root
# USER soapuser

# La aplicación corre como daemon con gocron scheduler
# Comando por defecto
CMD ["./noc-monitoring"]