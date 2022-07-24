package server

import (
	pb "ais-qa-assesment/gen"
	"ais-qa-assesment/models"
	"context"
	"log"
)

func (s *Server) GetTransaction(ctx context.Context, req *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {
	var response pb.GetTransactionResponse

	transaction, err := s.repository.GetTransaction(nil, req.GetId())
	if err != nil {
		log.Println("err-get-transaction: ", err)
		return nil, err
	}

	response.Transaction = models.TransactionToGrpcMessage(transaction)

	return &response, nil
}
