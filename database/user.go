package database

type UserBaseData struct {
	Username string `json:"username" example:"tranquanthuan"`
	Email    string `json:"email" example:"tranquangthuan132@gmail.com"`
	Password string `json:"password" example:"123456"`
	Phone    string `json:"phone" example:"0987654321"`
	Address  string `json:"address" example:"Ha Noi"`
	Birthday string `json:"birthday" example:"2005-08-15T15:52:01+00:00"`
	JoinedAt string `json:"joined_at" example:"2005-08-15T15:52:01+00:00"`
	Role     string `json:"role" example:"admin"`
	Rank     string `json:"rank" example:"1"`
	Gift     string `json:"gift" example:"0"`
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
	ID       string `json:"id" example:"tranquanthuan"`
	Username string `json:"username" example:"tranquanthuan"`
	Email    string `json:"email" example:"tranquangthuan132@gmail.com"`
	Password string `json:"password" example:"123456"`
	Phone    string `json:"phone" example:"0987654321"`
	Address  string `json:"address" example:"Ha Noi"`
	Birthday string `json:"birthday" example:"2005-08-15T15:52:01+00:00"`
	JoinedAt string `json:"joined_at" example:"2005-08-15T15:52:01+00:00"`
	Role     string `json:"role" example:"admin"`
	Rank     string `json:"rank" example:"1"`
	Gift     string `json:"gift" example:"0"`
	BaseModel
}
type LoginResponse struct {
	Code   int          `json:"code" example:"200"`
	Token  string       `json:"token" example:"iuniu32neui3rn38fh784e5yn78f5r57R&FGU*^TU?;.'grteuiHIUN98"`
	Expire string       `json:"expire" example:"2005-08-15T15:52:01+00:00"`
	Data   UserResponse `json:"data"`
}

func (UserResponse) TableName() string {
	return DB_ECOMMERCE + ".users"
}
