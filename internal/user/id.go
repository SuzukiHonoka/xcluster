package user

import (
	"errors"
	"xcluster/internal/database"
)

type ID uint

func (id ID) GetUser() (*User, error) {
	if id == 0 {
		return nil, errors.New("uid 0 is not allowed")
	}
	var user User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (id ID) DeleteUser() error {
	return database.DB.Delete(&User{}, id).Error
}
