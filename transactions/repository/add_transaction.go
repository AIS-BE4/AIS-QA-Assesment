package repository

import (
	"ais-qa-assesment/models"
	"log"

	"gorm.io/gorm"
)

func (r *Repository) AddTransaction(tx *gorm.DB, transaction *models.Transaction) error {
	db := r.db
	if tx != nil {
		db = tx
	}
	err := db.Create(transaction).Error
	if err != nil {
		log.Println("err-add-transaction: ", err)
		return err
	}
	return nil
}
