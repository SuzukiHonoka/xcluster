package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func newDatabaseConnection(config *Config, selectDatabase bool) (*gorm.DB, error) {
	dsn := config.generateDSN(selectDatabase)
	dialer := mysql.Open(dsn)
	db, err := gorm.Open(dialer, &gorm.Config{
		// warning, disable logger for production
		Logger: logger.Discard,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createDatabaseIfNotExist(db *gorm.DB, dbName string) error {
	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s ;", dbName)
	return db.Exec(sql).Error
}
