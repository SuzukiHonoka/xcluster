package user

import (
	"xcluster/internal/database"
)

type Name string

func (n Name) GetUser() (*User, error) {
	var user User
	if err := database.DB.First(&user, "user_name = ?", n).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
