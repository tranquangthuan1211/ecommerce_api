package database

type FeedBack struct {
	Id_User string `json:"id_user" example:"1"`
	Content string `json:"content" example:"1"`
	Rating  int    `json:"rating" example:"1"`
	BaseModel
}

type FeedBackResponse struct {
	Id string `json:"id" example:"1"`
	FeedBack
}

func (FeedBackResponse) TableName() string {
	return DB_ECOMMERCE + ".feedbacks"
}
