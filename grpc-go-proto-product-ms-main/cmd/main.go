package main

import (
    "fmt"
    "log"
    "net"

    "grpc-go-proto-product-ms-main/pkg/config"
    "grpc-go-proto-product-ms-main/pkg/db"
    "grpc-go-proto-product-ms-main/pkg/models"
    pb "grpc-go-proto-product-ms-main/pkg/proto"
    services "grpc-go-proto-product-ms-main/pkg/services"

    "google.golang.org/grpc"
)

func main() {
    // Cargar la configuración desde el archivo .env
    c, err := config.LoadConfig()
    if err != nil {
        log.Fatalln("Failed to load config:", err)
    }

    fmt.Println("Config:", c)

    // Inicializar la base de datos con los parámetros de configuración
    h := db.Init(c.DBHost, c.DBPort, c.DBUser, c.DBPass, c.DBName)

    // Ejecuta la migración
    h.DB.AutoMigrate(&models.Product{})
    h.DB.AutoMigrate(&models.StockDecreaseLog{})

    // Crear el listener
    lis, err := net.Listen("tcp", ":"+c.Port)
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    fmt.Println("Product Service on port", c.Port)

    // Crear un nuevo servidor gRPC
    grpcServer := grpc.NewServer()

    // Registrar el servicio de productos en el servidor gRPC
    pb.RegisterProductServiceServer(grpcServer, &services.Server{H: h})

    // Empezar a escuchar conexiones
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
