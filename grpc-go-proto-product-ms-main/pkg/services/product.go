package services

import (
	"context"
	"net/http"

	"grpc-go-proto-product-ms-main/pkg/db"
	"grpc-go-proto-product-ms-main/pkg/models"
	pb "grpc-go-proto-product-ms-main/pkg/pb"
)

type Server struct {
	pb.UnimplementedProductServiceServer
	H db.Handler
}

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	var product models.Product

	product.Name = req.Name
	product.Sku = req.Sku
	product.Category = req.Category
	product.Description = req.Description
	product.Stock = req.Stock
	product.Price = req.Price

	if result := s.H.DB.Create(&product); result.Error != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusConflict,
			Error:  []string{result.Error.Error()},
		}, nil
	}

	return &pb.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     product.Id,
	}, nil

}

func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var product models.Product

	if result := s.H.DB.First(&product, req.Id); result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  []string{result.Error.Error()},
		}, nil
	}

	data := &pb.FindOneData{
		Id:          product.Id,
		Name:        product.Name,
		Sku:         product.Sku,
		Category:    product.Category,
		Description: product.Description,
		Stock:       product.Stock,
		Price:       product.Price,
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (s *Server) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	var product models.Product

	if result := s.H.DB.First(&product, req.Id); result.Error != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusNotFound,
			Error:  []string{result.Error.Error()},
		}, nil
	}

	if product.Stock <= 0 {
		return &pb.DecreaseStockResponse{
			Status: http.StatusBadRequest,
			Error:  []string{"Product out of stock"},
		}, nil

	}

	var log models.StockDecreaseLog

	if result := s.H.DB.Where(&models.StockDecreaseLog{OrderId: req.OrderId}).First(&log); result.Error != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  []string{"Stock alredy decreased"},
		}, nil
	}

	product.Stock = product.Stock - 1

	s.H.DB.Save(&product)

	log.OrderId = req.OrderId
	log.Product = req.Id

	s.H.DB.Create(&log)

	return &pb.DecreaseStockResponse{
		Status: http.StatusOK,
	}, nil
}
