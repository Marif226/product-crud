package service

import (
	"log"

	"github.com/Marif226/product-crud/internal/repository"
)

type PurchaseServiceImpl struct {
	repo repository.PurchaseRepo
}

func NewPurchaseService(purchaseRepo repository.PurchaseRepo) *PurchaseServiceImpl {
	return &PurchaseServiceImpl{
		purchaseRepo,
	}
}

func (s *PurchaseServiceImpl) CreatePurchase() {
	s.repo.CreatePurchase()
	log.Println("Service Create Purchase")
}

func (s *PurchaseServiceImpl) GetPurchase() {
	s.repo.GetPurchase()
	log.Println("Service Get Purchase")
}

func (s *PurchaseServiceImpl) UpdatePurchase() {
	s.repo.UpdatePurchase()
	log.Println("Service Update Purchase")
}

func (s *PurchaseServiceImpl) DeletePurchase() {
	s.repo.DeletePurchase()
	log.Println("Service Delete Purchase")
}