package repository

import (
	"ais-qa-assesment/models"
	"ais-qa-assesment/utils/errors"
	"log"

	"gorm.io/gorm"
)

func (r *Repository) GetStoreByProductId(tx *gorm.DB, productId uint64) (*models.Store, error) {
	var result models.Product

	db := r.db
	if tx != nil {
		db = tx
	}

	err := db.
		Where("id = ?", productId).
		Preload("Store").
		Find(&result).Error

	if err == gorm.ErrRecordNotFound {
		log.Println("err-get-store-by-product-id: ", err)
		return nil, errors.ErrNotFound
	}
	if err != nil {
		log.Println("err-get-store-by-product-id: ", err)
		return nil, err
	}

	return result.Store, nil
}
