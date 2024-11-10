package database

type Order struct {
	Id_user    string `json:"id_user" example:"1"`
	TotalPrice int    `json:"total_price" example:"100"`
	BaseModel
}

type OrderResponse struct {
	ID string `json:"id" gorm:"column:id"`
	Order
}

type OrderUpdate struct {
	Id_user    string `json:"id_user" example:"1"`
	TotalPrice int    `json:"total_price" example:"100"`
	BaseModel
}

func (OrderResponse) TableName() string {
	return DB_ECOMMERCE + ".orders"
}
