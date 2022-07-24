package server

import (
	pb "ais-qa-assesment/gen"
	"ais-qa-assesment/models"
	"context"
	"log"
)

func (s *Server) ListTransactions(ctx context.Context, req *pb.ListTransactionsRequest) (*pb.ListTransactionsResponse, error) {
	var response pb.ListTransactionsResponse

	transactions, err := s.repository.ListTransactions(nil)
	if err != nil {
		log.Println("err-list-transactions: ", err)
		return nil, err
	}

	for _, transaction := range transactions {
		response.Transactions = append(response.Transactions, models.TransactionToGrpcMessage(&transaction))
	}

	return &response, nil
}
