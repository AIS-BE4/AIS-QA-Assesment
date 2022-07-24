package server

import (
	pb "ais-qa-assesment/gen"
	"ais-qa-assesment/models"
	"context"
	"log"
)

func (s *Server) ListStores(ctx context.Context, req *pb.ListStoresRequest) (*pb.ListStoresResponse, error) {
	var response pb.ListStoresResponse

	stores, err := s.repository.ListStores(nil)
	if err != nil {
		log.Println("err-list-stores: ", err)
		return nil, err
	}

	for _, store := range stores {
		response.Stores = append(response.Stores, models.StoreToGrpcMessage(&store))
	}

	return &response, nil
}
