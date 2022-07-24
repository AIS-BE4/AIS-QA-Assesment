package server

import (
	pb "ais-qa-assesment/gen"
	"ais-qa-assesment/models"
	"context"
	"log"
)

func (s *Server) GetStore(ctx context.Context, req *pb.GetStoreRequest) (*pb.GetStoreResponse, error) {
	var response pb.GetStoreResponse

	store, err := s.repository.GetStore(nil, req.GetId())
	if err != nil {
		log.Println("err-get-store: ", err)
		return nil, err
	}

	response.Store = models.StoreToGrpcMessage(store)

	return &response, nil
}
