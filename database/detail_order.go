package database

type DetailOrder struct {
	Id_Order   string `json:"id_order" example:"1"`
	Id_Product string `json:"id_product" example:"1"`
	Quantity   int    `json:"quantity" example:"1"`
	Id_Sale    string `json:"id_sale" example:"1"`
	BaseModel
}

func (DetailOrder) TableName() string {
	return DB_ECOMMERCE + ".detail_orders"
}
