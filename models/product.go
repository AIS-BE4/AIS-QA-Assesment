package models

import (
	pb "ais-qa-assesment/gen"
)

type Product struct {
	ID           uint64        `json:"id"`
	StoreID      uint64        `json:"store_id"`
	ProductName  string        `json:"product_name"`
	Quantity     string        `json:"quantity"`
	Price        string        `json:"price"`
	CreatedAt    int64         `json:"created_at"`
	UpdatedAt    int64         `json:"updated_at,omitempty"`
	Store        *Store        `json:"store" gorm:"foreignKey:store_id"`
	Transactions []Transaction `json:"transactions" gorm:"foreignKey:product_id"`
}

func ProductToGrpcMessage(product *Product) *pb.Product {
	return &pb.Product{
		Id:          product.ID,
		StoreId:     product.StoreID,
		ProductName: product.ProductName,
		Quantity:    product.Quantity,
		Price:       product.Price,
	}
}
