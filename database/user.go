package database

type UserBaseData struct {
	Username string
	Email    string
	Password string
	Phone    string
	Address  string
	Role     string
	BaseModel
}
type UserResponse struct {
	ID string `json:"id" gorm:"column:id"`
	UserBaseData
}
type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type Register struct {
	UserBaseData
}
type LoginResponse struct {
	Code   int          `json:"code" example:"200"`
	Token  string       `json:"token" example:"iuniu32neui3rn38fh784e5yn78f5r57R&FGU*^TU?;.'grteuiHIUN98"`
	Expire string       `json:"expire" example:"2005-08-15T15:52:01+00:00"`
	Data   UserResponse `json:"data"`
}

func (UserResponse) TableName() string {
	return DB_ECOMMERCE + ".USERS"
}
