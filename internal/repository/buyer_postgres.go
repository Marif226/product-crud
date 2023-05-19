package repository

import (
	"database/sql"
	"fmt"
	"github.com/Marif226/product-crud/internal/model"
)

// struct that implements BuyerRepo interface
type BuyerRepoImpl struct {
	db *sql.DB
}

func NewBuyerPostgresRepo(db *sql.DB) *BuyerRepoImpl {
	return &BuyerRepoImpl{
		db: db,
	}
}

// CreateBuyer creates a buyer object and returns its id. Returns 0 and error if failed.
func (r *BuyerRepoImpl) CreateBuyer(buyer model.Buyer) (int, error) {
	// log.Println("Repo Create Buyer")
	query := fmt.Sprintf("INSERT INTO %s (name, contact) values ($1, $2) RETURNING id;", buyersTable)

	row := r.db.QueryRow(query, buyer.Name, buyer.Contact)

	var id int
	err := row.Scan(&id)
	// log.Println(id)
	if err != nil {
		return 0, err
	}
	
	return id, nil
}

// GetAllBuyers returns the list of all existing buyers in the database
func (r *BuyerRepoImpl) GetAllBuyers() {

}

// GetBuyerById return buyer object with matched id with error if failed
func (r *BuyerRepoImpl) GetBuyerById(id int) (model.Buyer, error) {
	// log.Println("Repo Get Buyer")

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1;", buyersTable)

	row := r.db.QueryRow(query, id)

	var buyer model.Buyer
	err := row.Scan(&buyer.ID, &buyer.Name, &buyer.Contact)
	// log.Println(buyer)

	return buyer, err
}

// UpdateBuyer updates buyer with matched id and returns updated object or error if failed
func (r *BuyerRepoImpl) UpdateBuyer(buyer model.Buyer) (model.Buyer, error) {
	// log.Println("Repo Update Buyer")

	// query to update buyer in buyers table
	query := fmt.Sprintf("UPDATE %s SET name = $1, contact = $2 WHERE id = $3;", buyersTable)

	r.db.Exec(query, buyer.Name, buyer.Contact, buyer.ID)

	query = fmt.Sprintf("SELECT * FROM %s WHERE id = $1;", buyersTable)

	row := r.db.QueryRow(query, buyer.ID)

	var updatedBuyer model.Buyer
	err := row.Scan(&updatedBuyer.ID, &updatedBuyer.Name, &updatedBuyer.Contact)
	
	return updatedBuyer, err
}

// DeleteBuyer deletes buyer with the given id and returns error if failed
func (r *BuyerRepoImpl) DeleteBuyer(id int) error {
	// log.Println("Repo Delete Buyer")

	// query to dekete buyer in buyers table
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1;", buyersTable)

	_, err := r.db.Exec(query, id)

	return err
} 