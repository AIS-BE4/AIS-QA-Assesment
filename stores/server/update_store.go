package server

import (
	pb "ais-qa-assesment/gen"
	"ais-qa-assesment/models"
	"context"
	"fmt"
	"log"
	"time"
)

func (s *Server) UpdateStore(ctx context.Context, req *pb.UpdateStoreRequest) (*pb.UpdateStoreResponse, error) {
	var response pb.UpdateStoreResponse

	if req == nil || req.Store == nil {
		return nil, fmt.Errorf("store is not provided")
	}

	store := &models.Store{
		ID:        req.Store.Id,
		StoreName: req.Store.StoreName,
		Phone:     req.Store.Phone,
		Email:     req.Store.Email,
		Address:   req.Store.Address,
		UpdatedAt: time.Now().Unix(),
	}

	err := s.repository.UpdateStore(nil, store)

	if err != nil {
		log.Println("err-update-store: ", err)
		return nil, err
	}

	response.Store = models.StoreToGrpcMessage(store)

	return &response, nil
}
