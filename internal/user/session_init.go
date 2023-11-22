package user

import "xcluster/internal/database"

func InitUserSessionTable() error {
	return database.DB.AutoMigrate(&Session{})
}
