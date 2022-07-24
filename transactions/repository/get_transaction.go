package repository

import (
	"ais-qa-assesment/models"
	"log"

	"gorm.io/gorm"
)

func (r *Repository) GetTransaction(tx *gorm.DB, transactionId uint64) (*models.Transaction, error) {
	var result *models.Transaction

	db := r.db
	if tx != nil {
		db = tx
	}
	err := db.
		Where("id = ?", transactionId).
		Find(&result).Error

	if err == gorm.ErrRecordNotFound {
		log.Println("err-get-transaction: ", err)
		return nil, nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("err-get-transaction: ", err)
		return nil, err
	}

	return result, nil
}
