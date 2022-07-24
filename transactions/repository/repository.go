package repository

import (
	"ais-qa-assesment/models"

	"gorm.io/gorm"
)

type RepoMethod interface {
	AddTransaction(tx *gorm.DB, transaction *models.Transaction) error
	GetProduct(tx *gorm.DB, productId uint64) (*models.Product, error)
	GetTransaction(tx *gorm.DB, transactionId uint64) (*models.Transaction, error)
	ListTransactions(tx *gorm.DB) ([]models.Transaction, error)
	UpdateProduct(tx *gorm.DB, product *models.Product) error
}

type Repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) RepoMethod {
	return &Repository{
		db: db,
	}
}
