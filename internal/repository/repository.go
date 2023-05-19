package repository

import (
	"database/sql"

	"github.com/Marif226/product-crud/internal/model"
)

type BuyerRepo interface {
	CreateBuyer(model.Buyer) (int, error)
	GetAllBuyers() ([]model.Buyer, error)
	GetBuyerById(id int) (model.Buyer, error)
	UpdateBuyer(model.Buyer) (error)
	DeleteBuyer(id int) error
}

type PurchaseRepo interface {
	CreatePurchase(model.Purchase) (int, error)
	GetAllPurchases() ([]model.Purchase, error)
	GetPurchaseById(id int) (model.Purchase, error)
	UpdatePurchase(model.Purchase) (error)
	DeletePurchase(id int) error
}

type Repository struct {
	BuyerRepo
	PurchaseRepo
}

func New(db *sql.DB) *Repository{
	return &Repository{
		BuyerRepo: NewBuyerPostgresRepo(db),
		PurchaseRepo: NewPurchasePostgresRepo(db),
	}
}