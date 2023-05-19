package repository

import (
	"database/sql"
	"fmt"
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
func (r *PurchaseRepoImpl) GetAllPurchases() ([]model.Purchase, error) {
	// query to get all purchases from the database
	query := fmt.Sprintf("SELECT * FROM %s;", purchasesTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	// count number of purchases in the table
	var purchasesNumber int
	query = fmt.Sprintf("SELECT COUNT(*) FROM %s;", purchasesTable)
	err = r.db.QueryRow(query).Scan(&purchasesNumber)
	if err != nil {
		return nil, err
	}

	purchasesList := make([]model.Purchase, 0, purchasesNumber)

	// append purchases from rows to purchases list
	for rows.Next() {
		var purchase model.Purchase

		err := rows.Scan(&purchase.ID, &purchase.Name, &purchase.Description, &purchase.Quantity, &purchase.Price, &purchase.BuyerID)
		if err != nil {
			return purchasesList, err
		}

		purchasesList = append(purchasesList, purchase)
	}

	return purchasesList, nil
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

// UpdatePurchase updates purchase using passed purchase object
func (r *PurchaseRepoImpl) UpdatePurchase(purchase model.Purchase) (error) {
	// log.Println("Repo Update Purhcase")

	// query to update buyer in buyers table
	query := fmt.Sprintf("UPDATE %s SET name = $1, description = $2, quantity = $3, price = $4, buyer_id = $5 WHERE id = $6;", purchasesTable)

	_, err := r.db.Exec(query, purchase.Name, purchase.Description, purchase.Quantity, purchase.Price, purchase.BuyerID, purchase.ID)

	// query = fmt.Sprintf("SELECT * FROM %s WHERE id = $1;", purchasesTable)

	// row := r.db.QueryRow(query, buyer.ID)

	// var updatedPurchase model.Buyer
	// err := row.Scan(&updatedBuyer.ID, &updatedBuyer.Name, &updatedBuyer.Contact)
	
	return err
}

// DeletePurchase deletes purchase with specified id from the database
func (r *PurchaseRepoImpl) DeletePurchase(id int) error {
	// log.Println("Repo Delete Purhcase")

	// query to delete purchase in purchases table
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1;", purchasesTable)

	_, err := r.db.Exec(query, id)

	return err
}