# Proyecto RESTOCK

Proyecto de restaurant utilizando microservicios con grpc.

## Requisitos

Antes de comenzar, asegúrate de tener instalado en tu sistema los siguientes programas:

- [Node.js](https://nodejs.org/)
- [Go](https://golang.org/)
- [protoc](https://github.com/protocolbuffers/protobuf)
- [Docker Desktop](https://www.docker.com/products/docker-desktop)
- Una cuenta en [Docker Hub](https://hub.docker.com/)

## Instalación de Dependencias

Para instalar las dependencias en la API y microservicios de Nest.js, ejecuta el siguiente comando en la raíz del proyecto:

```bash
npm install
```

Para instalar dependecias de go:

```bash
go mod download
```

## Docker

Las imágenes de Docker se construyen automáticamente en el pipeline. Después de realizar un commit, las imágenes estarán disponibles y podrás iniciar el contenedor del proyecto localmente con el siguiente comando:

```bash
docker-compose up -d
```

Si realizas cambios en el código, asegúrate de eliminar el contenedor existente y las imágenes (excepto la imagen de PostgreSQL) antes de volver a iniciar el contenedor:

```bash
docker-compose up -d
```

## Actualización de estructura de Proto

Si realizas cambios en la estructura de los modelos de proto del microservicio de producto o directamente en el microservicio, ejecuta el siguiente comando para actualizar la estructura de proto en el microservicio:

```bash
protoc --go_out=../grpc-go-proto-product-ms-main/pkg --go_opt=paths=source_relative --go-grpc_out=../grpc-go-proto-product-ms-main/pkg --go-grpc_opt=paths=source_relative proto/product.proto
```

Recuerda ejecutar este comando después de realizar cambios en la estructura de proto.