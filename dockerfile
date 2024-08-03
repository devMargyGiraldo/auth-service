# Usa una imagen base oficial de Go para construir el binario
FROM golang:1.20-alpine AS builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos go.mod y go.sum para instalar las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto del código de la aplicación
COPY . .

# Compila la aplicación
RUN go build -o /auth_service

# Crea una imagen ligera para ejecutar el binario
FROM alpine:latest

# Establece un directorio de trabajo
WORKDIR /root/

# Copia el binario desde la imagen builder
COPY --from=builder /auth_service .

# Expone el puerto en el que correrá la aplicación
EXPOSE 8080

# Define el comando para ejecutar la aplicación
CMD ["./auth_service"]
