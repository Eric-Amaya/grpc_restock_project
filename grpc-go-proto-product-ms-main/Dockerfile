# Usar una imagen base oficial de Go para construir el binario
FROM golang:latest

RUN apt-get update \
 && DEBIAN_FRONTEND=noninteractive \
    apt-get install --no-install-recommends --assume-yes \
      protobuf-compiler

# Instalar los plugins de Go para protoc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
 && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos del proyecto al directorio de trabajo
COPY . .

# Descargar las dependencias
RUN go mod download

# Generar los archivos .pb.go y .pb.go.grpc a partir de los archivos .proto
#RUN protoc --go_out=..\grpc-go-proto-product-ms-main\pkg --go_opt=paths=source_relative --go-grpc_out=..\grpc-go-proto-product-ms-main\pkg --go-grpc_opt=paths=source_relative proto\product.proto

# Compilar el binario
RUN go build -o main ./cmd

# Exponer el puerto en el que corre el servicio
EXPOSE 5053

# Comando para ejecutar el binario
CMD ["./main"]
