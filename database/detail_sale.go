package database

type DeatailSale struct {
	SaleID         string `json:"sale_id" example:"1"`
	ProductID      string `json:"product_id" example:"1"`
	ExpirationDate string `json:"expiration_date" example:"2021-12-31"`
}

type DetailSaleResponse struct {
	ID string `json:"id" gorm:"column:id"`
	DeatailSale
}

func (DetailSaleResponse) TableName() string {
	return DB_ECOMMERCE + ".detail_sales"
}
