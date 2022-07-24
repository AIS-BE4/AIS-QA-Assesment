package repository

import (
	"ais-qa-assesment/models"
	"log"

	"gorm.io/gorm"
)

func (r *Repository) ListProducts(tx *gorm.DB) ([]models.Product, error) {
	var results []models.Product

	db := r.db
	if tx != nil {
		db = tx
	}
	err := db.
		Find(&results).Error

	if err == gorm.ErrRecordNotFound {
		log.Println("err-list-products: ", err)
		return nil, nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("err-list-products: ", err)
		return nil, err
	}

	return results, nil
}
