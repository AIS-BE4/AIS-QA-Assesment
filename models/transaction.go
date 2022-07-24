package models

import pb "ais-qa-assesment/gen"

type Transaction struct {
	ID        uint64   `json:"id"`
	ProductID uint64   `json:"product_id"`
	Quantity  string   `json:"quantity"`
	Total     string   `json:"total"`
	CreatedAt int64    `json:"created_at"`
	Product   *Product `json:"product" gorm:"foreignKey:product_id"`
}

func TransactionToGrpcMessage(transaction *Transaction) *pb.Transaction {
	return &pb.Transaction{
		Id:        transaction.ID,
		ProductId: transaction.ProductID,
		Quantity:  transaction.Quantity,
		Total:     transaction.Total,
	}
}
