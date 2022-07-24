package repository

import (
	"ais-qa-assesment/models"
	"ais-qa-assesment/utils/errors"
	"log"

	"gorm.io/gorm"
)

func (r *Repository) GetProductByTransactionId(tx *gorm.DB, transactionId uint64) (*models.Product, error) {
	var result models.Transaction

	db := r.db
	if tx != nil {
		db = tx
	}
	err := db.
		Where("id = ?", transactionId).
		Preload("Product").
		Find(&result).Error

	if err == gorm.ErrRecordNotFound {
		log.Println("err-get-product-by-transaction-id: ", err)
		return nil, errors.ErrNotFound
	}
	if err != nil {
		log.Println("err-get-product-by-transaction-id: ", err)
		return nil, err
	}

	return result.Product, nil
}
