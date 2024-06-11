# Usar una imagen base oficial de Go para construir el binario
FROM golang:latest

RUN apt-get update \
 && DEBIAN_FRONTEND=noninteractive \
    apt-get install --no-install-recommends --assume-yes \
      protobuf-compiler

# Instalar los plugins de Go para protoc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
 && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Agregar el directorio de binarios de Go al PATH
ENV PATH $PATH:/go/bin

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Eliminar la carpeta pkg/pb y su contenido para windows Remove-Item -Recurse -Force .\pkg\pb
RUN rm -rf pkg/pb

# Clonar el repositorio
RUN git clone https://github.com/Eric-Amaya/grpc-nest-proto-order-manager-app-main.git pkg/pb/

# Ejecutar el comando protoc
RUN protoc --proto_path=pkg/pb/proto --go_out=pkg/proto --go_opt=paths=source_relative --go-grpc_out=pkg/proto --go-grpc_opt=paths=source_relative pkg/pb/proto/product.proto

# Copiar los archivos del proyecto al directorio de trabajo
COPY . .

# Descargar las dependencias
RUN go mod download

# Compilar el binario
RUN go build -o main ./cmd

# Exponer el puerto en el que corre el servicio 5053
EXPOSE 5053

# Comando para ejecutar el binario
CMD ["./main"]
