package database

type Product struct {
	CategoryID        string `json:"category_id" example:"1"`
	Name              string `json:"name" example:"Ring"`
	Price             int    `json:"price" example:"100"`
	Manufacturer      string `json:"manufacturer" example:"Rolex"`
	Sale              int    `json:"sale" example:"10"`
	Max_quantity      int    `json:"max_quantity" example:"100"`
	Currency_quantity string `json:"currency" example:"USD"`
	BaseModel
}

type ProductResponse struct {
	ID string `json:"id" gorm:"column:id"`
	Product
}

type ProductUpdate struct {
	CategoryID        string `json:"category_id" example:"1"`
	Name              string `json:"name" example:"Ring"`
	Price             int    `json:"price" example:"100"`
	Manufacturer      string `json:"manufacturer" example:"Rolex"`
	Sale              int    `json:"sale" example:"10"`
	Max_quantity      int    `json:"max_quantity" example:"100"`
	Currency_quantity string `json:"currency" example:"USD"`
	BaseModel
}

func (ProductResponse) TableName() string {
	return DB_ECOMMERCE + ".products"
}
