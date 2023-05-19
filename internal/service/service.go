package service

import (
	"github.com/Marif226/product-crud/internal/model"
	"github.com/Marif226/product-crud/internal/repository"
)

type BuyerService interface {
	CreateBuyer(buyer model.Buyer) (int, error)
	GetAllBuyers()
	GetBuyerById(id int) (model.Buyer, error)
	UpdateBuyer(buyer model.Buyer) (model.Buyer, error)
	DeleteBuyer(id int) error
}

type PurchaseService interface {
	CreatePurchase(purchase model.Purchase) (int, error)
	GetAllPurchases()
	GetPurchaseById(id int) (model.PurchaseResponse, error)
	UpdatePurchase(purchase model.Purchase) (model.PurchaseResponse, error)
	DeletePurchase(id int) error
}

type Service struct {
	PurchaseService
	BuyerService
}

func New(repos *repository.Repository) *Service {
	return &Service{
		BuyerService: NewBuyerService(repos.BuyerRepo),
		PurchaseService: NewPurchaseService(repos.PurchaseRepo, repos.BuyerRepo),
	}
}