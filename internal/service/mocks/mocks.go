package mocks_service

import (
	"github.com/Marif226/product-crud/internal/model"
	"github.com/Marif226/product-crud/internal/repository"
	"github.com/Marif226/product-crud/internal/service"
)

func NewMockService(mr *repository.Repository) *service.Service {
	return &service.Service{
		BuyerService: &MockBuyerService{
			repo: mr.BuyerRepo,
		},
		PurchaseService: &MockPurchaseService{
			repo: mr.PurchaseRepo,
			buyerRepo: mr.BuyerRepo,
		},
	}
}

type MockBuyerService struct {
	repo repository.BuyerRepo
}

func (mb *MockBuyerService) CreateBuyer(buyer model.Buyer) (int, error) {
	return mb.repo.CreateBuyer(buyer)
}

func (mb *MockBuyerService) GetAllBuyers() ([]model.Buyer, error) {
	return mb.repo.GetAllBuyers()
}

func (mb *MockBuyerService) GetBuyerById(id int) (model.Buyer, error) {
	return mb.repo.GetBuyerById(id)
}

func (mb *MockBuyerService) UpdateBuyer(buyer model.Buyer) (model.Buyer, error) {
	var updatedBuyer model.Buyer
	return updatedBuyer, mb.repo.UpdateBuyer(buyer)
}

func (mb *MockBuyerService) DeleteBuyer(id int) (error) {
	return mb.repo.DeleteBuyer(id)
}

type MockPurchaseService struct {
	repo repository.PurchaseRepo
	buyerRepo repository.BuyerRepo
}

func (mp *MockPurchaseService) CreatePurchase(purchase model.Purchase) (int, error) {
	return mp.repo.CreatePurchase(purchase)
}

func (mp *MockPurchaseService) GetAllPurchases() ([]model.PurchaseResponse, error) {
	_, err := mp.repo.GetAllPurchases()
	if err != nil {
		return nil, err
	}

	var purchaseResponseList []model.PurchaseResponse

	return purchaseResponseList, nil
}

func (mp *MockPurchaseService) GetPurchaseById(id int) (model.PurchaseResponse, error) {
	var purchaseResponse model.PurchaseResponse
	_, err := mp.repo.GetPurchaseById(id)
	if err != nil {
		return purchaseResponse, err
	}
	return purchaseResponse, nil
}

func (mp *MockPurchaseService) UpdatePurchase(purchase model.Purchase) (model.PurchaseResponse, error) {
	var updatedPurchaseResponse model.PurchaseResponse
	err := mp.repo.UpdatePurchase(purchase)
	if err != nil {
		return updatedPurchaseResponse, err
	}
	return updatedPurchaseResponse, nil
}

func (mp *MockPurchaseService) DeletePurchase(id int) error {
	return mp.DeletePurchase(id)
}