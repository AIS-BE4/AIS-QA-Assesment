package models

import (
	pb "ais-qa-assesment/gen"
)

type Store struct {
	ID        uint64    `json:"id"`
	StoreName string    `json:"store_name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at,omitempty"`
	Products  []Product `json:"products" gorm:"foreignKey:store_id"`
}

func StoreToGrpcMessage(store *Store) *pb.Store {
	return &pb.Store{
		Id:        store.ID,
		StoreName: store.StoreName,
		Phone:     store.Phone,
		Email:     store.Email,
		Address:   store.Address,
	}
}
