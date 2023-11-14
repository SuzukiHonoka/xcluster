package user

import (
	"xcluster/internal/database"
)

type Name string

func (n Name) GetUser() (User, error) {
	var user User
	if err := database.DB.First(&user, "name = ?", n).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
