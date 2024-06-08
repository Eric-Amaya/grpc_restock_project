# Usar una imagen base oficial de Go para construir el binario
FROM golang:latest

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos del proyecto al directorio de trabajo
COPY . .

# Descargar las dependencias
RUN go mod download

# Compilar el binario
RUN go build -o main ./cmd

# Exponer el puerto en el que corre el servicio
EXPOSE 5053

# Comando para ejecutar el binario
CMD ["./main"]
