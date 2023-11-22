package server

import "xcluster/internal/database"

func InitGroupTable() error {
	return database.DB.AutoMigrate(&Group{})
}
