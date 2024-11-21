package database

type Sale struct {
	ID       string `json:"id" gorm:"column:id"`
	Name     string `json:"name" example:"Summer Sale"`
	Discount int    `json:"discount" example:"10"`
}

type SaleResponse struct {
	ID string `json:"id" gorm:"column:id"`
	Sale
}

type SaleUpdate struct {
	Name     string `json:"name" example:"Summer Sale"`
	Discount int    `json:"discount" example:"10"`
}

func (SaleResponse) TableName() string {
	return DB_ECOMMERCE + ".sales"
}
