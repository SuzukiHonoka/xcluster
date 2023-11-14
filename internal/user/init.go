package user

import "xcluster/internal/database"

func InitUserTable() error {
	return database.DB.AutoMigrate(User{})
}
