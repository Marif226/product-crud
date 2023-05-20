package mocks_repo

import (
	"errors"

	"github.com/Marif226/product-crud/internal/model"
	"github.com/Marif226/product-crud/internal/repository"
)

type MockRepository struct {
	MockBuyerRepo
	MockPurchaseRepo
}

func NewMockRepository() *repository.Repository {
	return &repository.Repository{
		BuyerRepo: &MockBuyerRepo{},
		PurchaseRepo: &MockPurchaseRepo{},
	}
}

type MockBuyerRepo struct {

}

func (mb *MockBuyerRepo) CreateBuyer(buyer model.Buyer) (int, error) {
	return 1, nil
}

func (mb *MockBuyerRepo) GetAllBuyers() ([]model.Buyer, error) {
	buyer, err := mb.GetBuyerById(1)
	if err != nil {
		return nil, err
	}

	var buyersList []model.Buyer

	buyersList = append(buyersList, buyer)

	return buyersList, nil
}

func (mb *MockBuyerRepo) GetBuyerById(id int) (model.Buyer, error) {
	var buyer model.Buyer
	if id != 1 {
		return buyer, errors.New("buyer with given id does not exist")
	}

	buyer = model.Buyer{
		ID: 1,
		Name: "Test Buyer",
		Contact: "test.buyer@example.com",
	}

	return buyer, nil
}

func (mb *MockBuyerRepo) UpdateBuyer(buyer model.Buyer) (error) {
	if buyer.ID != 1 {
		return errors.New("buyer with given does not exist")
	}

	return nil
}

func (mb *MockBuyerRepo) DeleteBuyer(id int) error {
	if id != 1 {
		return errors.New("buyer with given id does not exist")
	}

	return nil
}

type MockPurchaseRepo struct {

}

func (mp *MockPurchaseRepo) CreatePurchase(purchase model.Purchase) (int, error) {
	return 1, nil
}

func (mp *MockPurchaseRepo) GetAllPurchases() ([]model.Purchase, error) {
	var purchasesList []model.Purchase

	purchase := model.Purchase{
		ID: 1,
		Name: "Test Purchase",
		Description: "Test Description",
		Quantity: 1,
		Price: 10,
		BuyerID: 1,
	}

	purchasesList = append(purchasesList, purchase)

	return purchasesList, nil
}

func (mp *MockPurchaseRepo) GetPurchaseById(id int) (model.Purchase, error) {
	var purchase model.Purchase

	if id != 1 {
		return purchase, errors.New("purchase with given id does not exist")
	}

	purchase = model.Purchase{
		ID: 1,
		Name: "Test Purchase",
		Description: "Test Description",
		Quantity: 1,
		Price: 10,
		BuyerID: 1,
	}

	return purchase, nil
}

func (mp *MockPurchaseRepo) UpdatePurchase(purchase model.Purchase) (error) {
	if purchase.ID != 1 {
		return errors.New("purchase with given id does not exist")
	}

	return nil
}

func (mp *MockPurchaseRepo) DeletePurchase(id int) error {
	if id != 1 {
		return errors.New("purchase with given id does not exist")
	}

	return nil
}