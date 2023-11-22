package server

import "xcluster/internal/database"

func InitServerTable() error {
	return database.DB.AutoMigrate(&Server{})
}
