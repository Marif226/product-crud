package repository

import (
	"database/sql"
	"log"
)

type BuyerRepoImpl struct {
	db *sql.DB
}

func NewBuyerPostgresRepo(db *sql.DB) *BuyerRepoImpl {
	return &BuyerRepoImpl{
		db: db,
	}
}

func (r *BuyerRepoImpl) CreateBuyer() {
	log.Println("Repo Create Buyer")
}

func (r *BuyerRepoImpl) GetBuyerById() {
	log.Println("Repo Get Buyer")
}

func (r *BuyerRepoImpl) UpdateBuyer() {
	log.Println("Repo Update Buyer")
}

func (r *BuyerRepoImpl) DeleteBuyer() {
	log.Println("Repo Delete Buyer")
}