package service

import (
	"github.com/Marif226/product-crud/internal/model"
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

func (s *BuyerServiceImpl) CreateBuyer(buyer model.Buyer) (int, error) {
	// log.Println("Service Create Buyer")

	return s.repo.CreateBuyer(buyer)
}

func (s *BuyerServiceImpl) GetBuyerById(id int) (model.Buyer, error) {
	// log.Println("Service Get Buyer By Id")

	return s.repo.GetBuyerById(id)
}

func (s *BuyerServiceImpl) UpdateBuyer(buyer model.Buyer) (model.Buyer, error) {
	// log.Println("Service Update Buyer")

	return s.repo.UpdateBuyer(buyer)
}

func (s *BuyerServiceImpl) DeleteBuyer(id int) error {
	return s.repo.DeleteBuyer(id)

	// log.Println("Service Delete Buyer")
}