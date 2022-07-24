package server

import (
	pb "ais-qa-assesment/gen"
	"ais-qa-assesment/models"
	"context"
	"fmt"
	"log"
	"time"
)

func (s *Server) AddStore(ctx context.Context, req *pb.AddStoreRequest) (*pb.AddStoreResponse, error) {
	var response pb.AddStoreResponse

	if req == nil || req.Store == nil {
		return nil, fmt.Errorf("store is not provided")
	}

	store := &models.Store{
		StoreName: req.Store.StoreName,
		Phone:     req.Store.Phone,
		Email:     req.Store.Email,
		Address:   req.Store.Address,
		CreatedAt: time.Now().Unix(),
	}

	err := s.repository.AddStore(nil, store)

	if err != nil {
		log.Println("err-add-store: ", err)
		return nil, err
	}

	response.Store = models.StoreToGrpcMessage(store)

	return &response, nil
}
