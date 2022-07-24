package repository

import (
	"ais-qa-assesment/models"
	"log"

	"gorm.io/gorm"
)

func (r *Repository) AddStore(tx *gorm.DB, store *models.Store) error {
	db := r.db
	if tx != nil {
		db = tx
	}
	err := db.Create(store).Error
	if err != nil {
		log.Println("err-add-store: ", err)
		return err
	}
	return nil
}
