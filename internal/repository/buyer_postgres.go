package repository

import (
	"log"
)

type BuyerRepoImpl struct {

}

func NewBuyerPostgresRepo() *BuyerRepoImpl {
	return &BuyerRepoImpl{
		
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