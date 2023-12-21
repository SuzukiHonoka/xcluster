package database

import (
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewDatabaseWrapper(config *Config, create bool) (*Database, error) {
	if create {
		// create database if not exist by using one-time connection
		db, err := newDatabaseConnection(config, false)
		if err != nil {
			return nil, err
		}
		if err = createDatabaseIfNotExist(db, config.DatabaseName); err != nil {
			return nil, err
		}
		// close one-time connection
		mdb, err := db.DB()
		if err != nil {
			return nil, err
		}
		if err = mdb.Close(); err != nil {
			return nil, err
		}
	}
	// actually get the database connection
	db, err := newDatabaseConnection(config, true)
	if err != nil {
		return nil, err
	}
	return &Database{
		DB: db,
	}, nil
}

func (d *Database) Close() error {
	db, err := d.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
