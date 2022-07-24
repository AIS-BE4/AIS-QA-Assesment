package repository

import (
	"ais-qa-assesment/models"

	"gorm.io/gorm"
)

type RepoMethod interface {
	AddStore(tx *gorm.DB, store *models.Store) error
	GetStore(tx *gorm.DB, storeId uint64) (*models.Store, error)
	GetStoreByProductId(tx *gorm.DB, productId uint64) (*models.Store, error)
	ListStores(tx *gorm.DB) ([]models.Store, error)
	UpdateStore(tx *gorm.DB, store *models.Store) error
}

type Repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) RepoMethod {
	return &Repository{
		db: db,
	}
}
