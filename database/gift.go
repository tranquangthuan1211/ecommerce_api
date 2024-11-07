package database

type Gift struct {
	Name string `json:"name" example:"DIAMOND"`
	Sale int    `json:"sale" example:"10"`
	BaseModel
}
type GiftResponse struct {
	ID string `json:"id" gorm:"column:id"`
	Gift
}
type GiftUpdate struct {
	Name string `json:"name" example:"DIAMOND"`
	Sale int    `json:"sale" example:"10"`
	BaseModel
}

func (GiftResponse) TableName() string {
	return DB_ECOMMERCE + ".gifts"
}
