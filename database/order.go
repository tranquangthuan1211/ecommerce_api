package database

type Order struct {
	Id_user    string `json:"id_user" example:"1"`
	Product_id string `json:"product_id" example:"1"`
	Quantity   int    `json:"quantity" example:"1"`
	TotalPrice int    `json:"total_price" example:"100"`
	BaseModel
}

type OrderResponse struct {
	ID string `json:"id" gorm:"column:id"`
	Order
}
type OrderDetail struct {
	OrderResponse
	SalesDate string `json:"sales_date" example:"2021-01-01"`
}
type OrderUpdate struct {
	Id_user    string `json:"id_user" example:"1"`
	TotalPrice int    `json:"total_price" example:"100"`
	BaseModel
}

func (OrderResponse) TableName() string {
	return DB_ECOMMERCE + ".orders"
}
