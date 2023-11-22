package user

import "xcluster/internal/database"

type Users []*User

func All() (Users, error) {
	var users Users
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
