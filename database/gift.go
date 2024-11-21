package database

type Gift struct {
	ID_RANK string `json:"id_rank" example:"1"`
	ID_USER string `json:"id_user" example:"1"`
	BaseModel
}
type GiftResponse struct {
	ID string `json:"id" gorm:"column:id"`
	Gift
}
type GiftDetail struct {
	Name string `json:"rank_name" example:"Gift 1"`
	Sale string `json:"sale" example:"10"`
	BaseModel
}
type GiftForUser struct {
	Name string `json:"rank_name" example:"Gift 1"`
	Sale string `json:"sale" example:"10"`
	BaseModel
}
type GiftUpdate struct {
	ID_RANK string `json:"id_rank" example:"1"`
	ID_USER string `json:"id_user" example:"1"`
	BaseModel
}

func (GiftResponse) TableName() string {
	return DB_ECOMMERCE + ".gifts"
}
