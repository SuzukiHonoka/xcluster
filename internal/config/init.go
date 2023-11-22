package config

import "xcluster/internal/database"

func InitConfigTable() error {
	return database.DB.AutoMigrate(&Config{})
}

func InitCache() {
	cache = NewCache()
}
