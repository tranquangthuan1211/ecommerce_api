package database

type Rank struct {
	Name string `json:"name" example:"DIAMOND"`
	BaseModel
}

type RankResponse struct {
	ID string `json:"id" gorm:"column:id"`
	Rank
}

type RankUpdate struct {
	Name string `json:"name" example:"DIAMOND"`
}

func (RankResponse) TableName() string {
	return DB_ECOMMERCE + ".ranks"
}
