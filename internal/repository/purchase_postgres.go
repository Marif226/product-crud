package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Marif226/product-crud/internal/model"
)

type PurchaseRepoImpl struct {
	db *sql.DB
}

func NewPurchasePostgresRepo(db *sql.DB) *PurchaseRepoImpl {
	return &PurchaseRepoImpl{
		db: db,
	}
}

// CreatePurchase creates a purchase object with specified parameters and returns its id. Returns 0 and error if failed.
func (r *PurchaseRepoImpl) CreatePurchase(purchase model.Purchase) (int, error) {
	// log.Println("Repo Create Purhcase")

	query := fmt.Sprintf("INSERT INTO %s (name, description, quantity, price, buyer_id) values ($1, $2, $3, $4, $5) RETURNING id;", purchasesTable)

	row := r.db.QueryRow(query, purchase.Name, purchase.Description, purchase.Quantity, purchase.Price, purchase.BuyerID)

	var id int
	err := row.Scan(&id)
	// log.Println(id)
	if err != nil {
		return 0, err
	}
	
	return id, nil
}

// GetAllPurchases returns the list of all existing purchases in the database
func (r *PurchaseRepoImpl) GetAllPurchases() {

}

// GetPurchase returns purchase object by specified id. Returns error if failed to find.
func (r *PurchaseRepoImpl) GetPurchaseById(id int) (model.Purchase, error) {
	// log.Println("Repo Get Purhcase")

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1;", purchasesTable)

	row := r.db.QueryRow(query, id)

	var purchase model.Purchase
	err := row.Scan(&purchase.ID, &purchase.Name, &purchase.Description, &purchase.Quantity, &purchase.Price, &purchase.BuyerID)
	// log.Println(buyer)

	return purchase, err
}

func (r *PurchaseRepoImpl) UpdatePurchase() {
	log.Println("Repo Update Purhcase")
}

func (r *PurchaseRepoImpl) DeletePurchase() {
	log.Println("Repo Delete Purhcase")
}