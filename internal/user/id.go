package user

import (
	"xcluster/internal/database"
)

type ID uint

func (id ID) GetUser() (*User, error) {
	var user User
	if err := database.DB.First(&user, "user_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (id ID) GetSessions() (Sessions, error) {
	var sessions Sessions
	if err := database.DB.Find(&sessions, "user_id = ?", id).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

func (id ID) DeleteSessions() error {
	return database.DB.Delete(&Session{}, "user_id = ?", id).Error
}

func (id ID) DeleteUser() error {
	return database.DB.Delete(&User{}, id).Error
}
