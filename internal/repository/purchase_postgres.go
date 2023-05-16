package repository

import "log"

type PurchaseRepoImpl struct {

}

func NewPurchasePostgresRepo() *PurchaseRepoImpl {
	return &PurchaseRepoImpl{
		
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