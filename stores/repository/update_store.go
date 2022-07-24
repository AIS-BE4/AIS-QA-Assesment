package repository

import (
	"ais-qa-assesment/models"
	"log"

	"gorm.io/gorm"
)

func (r *Repository) UpdateStore(tx *gorm.DB, store *models.Store) error {
	db := r.db
	if tx != nil {
		db = tx
	}
	err := db.
		Updates(&store).Error

	if err == gorm.ErrRecordNotFound {
		log.Println("err-update-store: ", err)
		return nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("err-update-store: ", err)
		return err
	}

	return nil
}
