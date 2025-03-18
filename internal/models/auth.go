package models

type Auth struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) (bool, error) {
	var auth Auth
	err := db.Select("id").Where(&Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
