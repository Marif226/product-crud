package repository

import "database/sql"

type BuyerRepo interface {
	CreateBuyer()
	GetBuyerById()
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