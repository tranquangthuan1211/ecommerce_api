package database

type Category struct {
	Name string `json:"name" example:"Ring"`
	BaseModel
}

type CategoryResponse struct {
	ID string `json:"id" gorm:"column:id"`
	Category
}

type CategoryUpdate struct {
	Name string `json:"name" example:"Ring"`
}

func (CategoryResponse) TableName() string {
	return DB_ECOMMERCE + ".categories"
}
