package service

import (
	"github.com/Marif226/product-crud/internal/model"
	"github.com/Marif226/product-crud/internal/repository"
)

type PurchaseServiceImpl struct {
	repo repository.PurchaseRepo
	buyerRepo repository.BuyerRepo
}

func NewPurchaseService(purchaseRepo repository.PurchaseRepo, buyerRepo repository.BuyerRepo) *PurchaseServiceImpl {
	return &PurchaseServiceImpl{
		purchaseRepo,
		buyerRepo,
	}
}

func (s *PurchaseServiceImpl) CreatePurchase(purchase model.Purchase) (int, error) {
	// check if buyer with given exists
	_, err := s.buyerRepo.GetBuyerById(purchase.BuyerID)
	if err != nil {
		return 0, nil
	}
	
	return s.repo.CreatePurchase(purchase)

	// log.Println("Service Create Purchase")
}

func (s *PurchaseServiceImpl) GetAllPurchases() ([]model.PurchaseResponse, error) {
	purchasesList, err := s.repo.GetAllPurchases()
	if err != nil {
		return nil, err
	}

	purchasesResponseList := make([]model.PurchaseResponse, 0, len(purchasesList))

	for _, purchase := range purchasesList {
		// find buyer of the purchase by id
		buyer, err := s.buyerRepo.GetBuyerById(purchase.BuyerID)
		if err != nil {
			return nil, err
		}

		purchaseResponse := model.PurchaseResponse{
			ID: purchase.ID,
			Name: purchase.Name,
			Description: purchase.Description,
			Quantity: purchase.Quantity,
			Price: purchase.Price,
			Buyer: buyer,
		}

		purchasesResponseList = append(purchasesResponseList, purchaseResponse)
	}

	return purchasesResponseList, nil
}

func (s *PurchaseServiceImpl) GetPurchaseById(id int) (model.PurchaseResponse, error) {
	var purchaseResponse model.PurchaseResponse

	purchase, err := s.repo.GetPurchaseById(id)
	if err != nil {
		return purchaseResponse, err
	}

	// find buyer of the purchase by id
	buyer, err := s.buyerRepo.GetBuyerById(purchase.ID)
	if err != nil {
		return purchaseResponse, err
	}

	// wrap purchase info into purchase response with buyer object
	purchaseResponse = model.PurchaseResponse{
		ID: purchase.ID,
		Name: purchase.Name,
		Description: purchase.Description,
		Quantity: purchase.Quantity,
		Price: purchase.Price,
		Buyer: buyer,
	}

	return purchaseResponse, nil

	// log.Println("Service Get Purchase")
}

func (s *PurchaseServiceImpl) UpdatePurchase(purchase model.Purchase) (model.PurchaseResponse, error) {
	var updatedPurchaseResponse model.PurchaseResponse

	err := s.repo.UpdatePurchase(purchase)
	if err != nil {
		return updatedPurchaseResponse, err
	}

	updatedPurchase, err := s.repo.GetPurchaseById(purchase.ID)
	if err != nil {
		return updatedPurchaseResponse, err
	}

	// find buyer of the purchase by id
	buyer, err := s.buyerRepo.GetBuyerById(purchase.ID)
	if err != nil {
		return updatedPurchaseResponse, err
	}

	// wrap purchase info into purchase response with buyer object
	updatedPurchaseResponse = model.PurchaseResponse{
		ID: purchase.ID,
		Name: updatedPurchase.Name,
		Description: updatedPurchase.Description,
		Quantity: updatedPurchase.Quantity,
		Price: updatedPurchase.Price,
		Buyer: buyer,
	}

	return updatedPurchaseResponse, nil

	// log.Println("Service Update Purchase")
}

func (s *PurchaseServiceImpl) DeletePurchase(id int) error {
	return s.repo.DeletePurchase(id)
	// log.Println("Service Delete Purchase")
}