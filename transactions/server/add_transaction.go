package server

import (
	pb "ais-qa-assesment/gen"
	"ais-qa-assesment/models"
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
)

func (s *Server) AddTransaction(ctx context.Context, req *pb.AddTransactionRequest) (*pb.AddTransactionResponse, error) {
	var response pb.AddTransactionResponse

	if req == nil || req.Transaction == nil {
		return nil, fmt.Errorf("transaction is not provided")
	}

	product, err := s.repository.GetProduct(nil, req.GetTransaction().ProductId)
	if err != nil {
		log.Println("err-get-product: ", err)
		return nil, err
	}

	transaction := &models.Transaction{
		ProductID: req.Transaction.ProductId,
		Quantity:  req.Transaction.Quantity,
		Total:     req.Transaction.Total,
		CreatedAt: time.Now().Unix(),
	}

	prodQuan, _ := strconv.Atoi(product.Quantity)
	tranQuan, _ := strconv.Atoi(transaction.Quantity)

	if prodQuan < tranQuan {
		log.Println("err-get-product: ", "unable to perform request as product quantity is insufficient")
		return nil, fmt.Errorf("unable to perform request as product quantity is insufficient")
	}

	tx := s.db.Begin()
	tx.Begin()
	{
		price, _ := strconv.Atoi(product.Price)
		finQuan := strconv.Itoa(prodQuan - tranQuan)
		transaction.Total = strconv.Itoa(price * tranQuan)

		product = &models.Product{
			ID:          product.ID,
			ProductName: product.ProductName,
			StoreID:     product.StoreID,
			Quantity:    finQuan,
			Price:       product.Price,
			UpdatedAt:   time.Now().Unix(),
		}

		err = s.repository.UpdateProduct(nil, product)
		if err != nil {
			log.Println("err-update-product: ", err)
			tx.Rollback()
		}

		err = s.repository.AddTransaction(nil, transaction)

		if err != nil {
			log.Println("err-add-transaction: ", err)
			tx.Rollback()
		}
		tx.Commit()
	}

	response.Transaction = models.TransactionToGrpcMessage(transaction)

	return &response, nil
}
