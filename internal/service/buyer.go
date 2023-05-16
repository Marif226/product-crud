package service

import (
	"log"

	"github.com/Marif226/product-crud/internal/repository"
)

type BuyerServiceImpl struct {
	repo repository.BuyerRepo
}

func NewBuyerService(buyerRepo repository.BuyerRepo) *BuyerServiceImpl {
	return &BuyerServiceImpl{
		buyerRepo,
	}
}

func (s *BuyerServiceImpl) CreateBuyer() {
	s.repo.CreateBuyer()
	log.Println("Service Create Buyer")
}

func (s *BuyerServiceImpl) GetBuyerById() {
	s.repo.GetBuyerById()
	log.Println("Service Get Buyer By Id")
}

func (s *BuyerServiceImpl) UpdateBuyer() {
	s.repo.UpdateBuyer()
	log.Println("Service Update Buyer")
}

func (s *BuyerServiceImpl) DeleteBuyer() {
	s.repo.DeleteBuyer()
	log.Println("Service Delete Buyer")
}