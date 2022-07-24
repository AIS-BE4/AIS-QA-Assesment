package server

import (
	pb "ais-qa-assesment/gen"
	"ais-qa-assesment/models"
	"context"
	"log"
)

func (s *Server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	var response pb.GetProductResponse

	product, err := s.repository.GetProduct(nil, req.GetId())
	if err != nil {
		log.Println("err-get-product: ", err)
		return nil, err
	}

	response.Product = models.ProductToGrpcMessage(product)

	return &response, nil
}
