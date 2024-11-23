package database

type CartDataBase struct {
	UserID    string `json:"user_id" gorm:"column:user_id"`
	ProductID string `json:"product_id" gorm:"column:product_id"`
	Quantity  int    `json:"quantity" gorm:"column:quantity"`
	BaseModel
}

type CartResponse struct {
	ID string `json:"id" gorm:"column:id"`
	BaseModel
}

func (CartResponse) TableName() string {
	return DB_ECOMMERCE + ".carts"
}
