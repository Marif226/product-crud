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

	// query to create buyer
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
func (r *BuyerRepoImpl) GetAllBuyers() ([]model.Buyer, error) {
	// query to get all buyers from the database
	query := fmt.Sprintf("SELECT * FROM %s;", buyersTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	// count number of buyers in the table
	var buyersNumber int
	query = fmt.Sprintf("SELECT COUNT(*) FROM %s;", buyersTable)
	err = r.db.QueryRow(query).Scan(&buyersNumber)
	if err != nil {
		return nil, err
	}

	buyersList := make([]model.Buyer, 0, buyersNumber)

	// append buyers from rows to buyers list
	for rows.Next() {
		var buyer model.Buyer

		err := rows.Scan(&buyer.ID, &buyer.Name, &buyer.Contact)
		if err != nil {
			return buyersList, err
		}

		buyersList = append(buyersList, buyer)
	}

	return buyersList, nil
}

// GetBuyerById return buyer object with matched id with error if failed
func (r *BuyerRepoImpl) GetBuyerById(id int) (model.Buyer, error) {
	// log.Println("Repo Get Buyer")

	// query to find buyer with specified id
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1;", buyersTable)

	row := r.db.QueryRow(query, id)

	var buyer model.Buyer
	err := row.Scan(&buyer.ID, &buyer.Name, &buyer.Contact)
	// log.Println(buyer)

	return buyer, err
}

// UpdateBuyer updates buyer with matched id and returns updated object or error if failed
func (r *BuyerRepoImpl) UpdateBuyer(buyer model.Buyer) (error) {
	// log.Println("Repo Update Buyer")

	// query to update buyer in buyers table
	query := fmt.Sprintf("UPDATE %s SET name = $1, contact = $2 WHERE id = $3;", buyersTable)

	_, err := r.db.Exec(query, buyer.Name, buyer.Contact, buyer.ID)

	// query = fmt.Sprintf("SELECT * FROM %s WHERE id = $1;", buyersTable)

	// row := r.db.QueryRow(query, buyer.ID)

	// var updatedBuyer model.Buyer
	// err := row.Scan(&updatedBuyer.ID, &updatedBuyer.Name, &updatedBuyer.Contact)
	
	return err
}

// DeleteBuyer deletes buyer with the given id and returns error if failed
func (r *BuyerRepoImpl) DeleteBuyer(id int) error {
	// log.Println("Repo Delete Buyer")

	// query to delete buyer in buyers table
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1;", buyersTable)

	_, err := r.db.Exec(query, id)

	return err
} 