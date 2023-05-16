package model

type Buyer struct {
	ID		int		`json:"id"`
	Name	string	`json:"name"`
	Contact	string	`json:"contact"`
}

type Purchase struct {
	ID			int		`json:"id"`
	Name		string	`json:"name"`
	Description	string	`json:"description"`
	Quantity	int		`json:"qunatity"`
	Price		int		`json:"price"`
	BuyerID		int		`json:"buyer_id"`
}