package user

import "xcluster/internal/database"

type Email string

func (n Email) GetUser() (*User, error) {
	var user User
	if err := database.DB.First(&user, "user_email = ?", n).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
