package repository

import (
	"ais-qa-assesment/models"
	"log"

	"gorm.io/gorm"
)

func (r *Repository) ListStores(tx *gorm.DB) ([]models.Store, error) {
	var results []models.Store

	db := r.db
	if tx != nil {
		db = tx
	}
	err := db.
		Find(&results).Error

	if err == gorm.ErrRecordNotFound {
		log.Println("err-list-stores: ", err)
		return nil, nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("err-list-stores: ", err)
		return nil, err
	}

	return results, nil
}
