package repository

import (
	"ais-qa-assesment/models"
	"log"

	"gorm.io/gorm"
)

func (r *Repository) AddProduct(tx *gorm.DB, product *models.Product) error {
	db := r.db
	if tx != nil {
		db = tx
	}
	err := db.Create(product).Error
	if err != nil {
		log.Println("err-add-product: ", err)
		return err
	}
	return nil
}
