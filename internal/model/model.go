package model

type Buyer struct {
	ID		int		`json:"id,string"`
	Name	string	`json:"name"`
	Contact	string	`json:"contact"`
}

type Purchase struct {
	ID			int		`json:"id"`
	Name		string	`json:"name"`
	Description	string	`json:"description"`
	Quantity	int		`json:"quantity,string"`
	Price		int		`json:"price,string"`
	BuyerID		int		`json:"buyer_id,string"`
}

type PurchaseResponse struct {
	ID			int
	Name		string
	Description	string
	Quantity	int
	Price		int
	Buyer		Buyer
}