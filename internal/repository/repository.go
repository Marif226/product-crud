package repository

type BuyerRepo interface {
	CreateBuyer()
	GetBuyerById()
	UpdateBuyer()
	DeleteBuyer()
}

type PurchaseRepo interface {
	CreatePurchase()
	GetPurchase()
	UpdatePurchase()
	DeletePurchase()
}

type Repository struct {
	BuyerRepo
	PurchaseRepo
}

func New() *Repository{
	return &Repository{
		BuyerRepo: NewBuyerPostgresRepo(),
		PurchaseRepo: NewPurchasePostgresRepo(),
	}
}