package server

import (
	pb "ais-qa-assesment/gen"
	"ais-qa-assesment/models"
	"context"
	"log"

	"gorm.io/gorm"
)

func (s *Server) GetStoreByProductId(ctx context.Context, req *pb.GetStoreByProductIdRequest) (*pb.GetStoreByProductIdResponse, error) {
	var response pb.GetStoreByProductIdResponse

	store, err := s.repository.GetStoreByProductId(nil, req.GetId())
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("err-get-store: ", err)
		return nil, err
	}
	if store != nil {
		response.Store = models.StoreToGrpcMessage(store)
	}

	return &response, nil
}
