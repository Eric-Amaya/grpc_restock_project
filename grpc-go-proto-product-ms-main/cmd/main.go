package main

import (
	"fmt"
	"log"
	"net"

	"grpc-go-proto-product-ms-main/pkg/config"
	"grpc-go-proto-product-ms-main/pkg/db"
	pb "grpc-go-proto-product-ms-main/pkg/pb"
	services "grpc-go-proto-product-ms-main/pkg/services"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	fmt.Println(c)

	// Inicializar la base de datos sin pasar ningún argumento
	h := db.Init()

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
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
