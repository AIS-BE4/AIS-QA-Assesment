package repository

import (
	"ais-qa-assesment/models"
	"log"

	"gorm.io/gorm"
)

func (r *Repository) GetStore(tx *gorm.DB, storeId uint64) (*models.Store, error) {
	var result *models.Store

	db := r.db
	if tx != nil {
		db = tx
	}
	err := db.
		Where("id = ?", storeId).
		Find(&result).Error

	if err == gorm.ErrRecordNotFound {
		log.Println("err-get-store: ", err)
		return nil, nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("err-get-store: ", err)
		return nil, err
	}

	return result, nil
}
