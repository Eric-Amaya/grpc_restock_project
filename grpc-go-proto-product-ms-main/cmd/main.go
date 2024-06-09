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
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	fmt.Println(c)

	h := db.Init(c.DBUrl)

    h.DB.AutoMigrate(&models.Product{})
    h.DB.AutoMigrate(&models.StockDecreaseLog{})

	address := fmt.Sprintf(":%s", c.Port)
    lis, err := net.Listen("tcp", address)
    if err != nil {
        log.Fatalln("Failed to listen:", err)
    }

    fmt.Println("Product Svc on", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

