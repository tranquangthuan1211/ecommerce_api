package database

type ClientBaseData struct {
	ClientName string `json:"client_name" example:"tranquanthuan"`
	Email      string `json:"email" example:"tranquangthuan132@gmail.com"`
	Password   string `json:"password" example:"123456"`
	Phone      string `json:"phone" example:"0987654321"`
	Address    string `json:"address" example:"Ha Noi"`
	Birthday   string `json:"birthday" example:"2005-08-15T15:52:01+00:00"`
	JoinedAt   string `json:"joined_at" example:"2005-08-15T15:52:01+00:00"`
	Rank       string `json:"rank" example:"1"`
	Gift       string `json:"gift" example:"0"`
	BaseModel
}
type ClientResponse struct {
	ID string `json:"id" gorm:"column:id"`
	ClientBaseData
}

func (ClientResponse) TableName() string {
	return DB_ECOMMERCE + ".client"
}
