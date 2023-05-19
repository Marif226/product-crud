package repository

import (
	"database/sql"

	"github.com/Marif226/product-crud/internal/model"
)

type BuyerRepo interface {
	CreateBuyer(model.Buyer) (int, error)
	GetBuyerById(id int) (model.Buyer, error)
	UpdateBuyer()
	DeleteBuyer()
}

type PurchaseRepo interface {
	CreatePurchase()
	GetPurchase()
	UpdatePurchase()
	DeletePurchase()
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