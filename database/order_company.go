package database

type OrderCompany struct {
	ID_COMPANY string `json:"id_company" example:"1"`
	QUANTITY   int    `json:"quantity" example:"1"`
	PRICE      int    `json:"price" example:"100000"`
	STATUS     string `json:"status" example:"pending"`
	BaseModel
}

type OrderCompanyResponse struct {
	ID string `json:"id" gorm:"column:id"`
	OrderCompany
}

func (OrderCompanyResponse) TableName() string {
	return DB_ECOMMERCE + ".order_companies"
}
