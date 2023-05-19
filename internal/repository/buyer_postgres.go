package repository

import (
	"database/sql"
	"log"
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
	log.Println("Repo Create Buyer")
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

func (r *BuyerRepoImpl) GetBuyerById(id int) (model.Buyer, error) {
	log.Println("Repo Get Buyer")

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1;", buyersTable)

	row := r.db.QueryRow(query, id)

	var buyer model.Buyer
	err := row.Scan(&buyer.ID, &buyer.Name, &buyer.Contact)
	log.Println(buyer)
	if err != nil {
		return buyer, err
	}

	return buyer, nil
}

func (r *BuyerRepoImpl) UpdateBuyer() {
	log.Println("Repo Update Buyer")
}

func (r *BuyerRepoImpl) DeleteBuyer() {
	log.Println("Repo Delete Buyer")
}