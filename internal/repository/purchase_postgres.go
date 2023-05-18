package repository

import (
	"database/sql"
	"log"
)

type PurchaseRepoImpl struct {
	db *sql.DB
}

func NewPurchasePostgresRepo(db *sql.DB) *PurchaseRepoImpl {
	return &PurchaseRepoImpl{
		db: db,
	}
}

func (r *PurchaseRepoImpl) CreatePurchase() {
	log.Println("Repo Create Purhcase")
}

func (r *PurchaseRepoImpl) GetPurchase() {
	log.Println("Repo Get Purhcase")
}

func (r *PurchaseRepoImpl) UpdatePurchase() {
	log.Println("Repo Update Purhcase")
}

func (r *PurchaseRepoImpl) DeletePurchase() {
	log.Println("Repo Delete Purhcase")
}