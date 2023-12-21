package server

import "xcluster/internal/database"

func InitTokenTable() error {
	return database.DB.AutoMigrate(&Token{})
}
