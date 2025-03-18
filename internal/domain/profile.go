package domain

type Profile struct {
	UserID      uint   `json:"user_id" gorm:"primary_key"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}