package database

import (
	"fmt"
	"testing"
)

var (
	databaseAddr     = "127.0.0.1:3306"
	databaseUser     = "root"
	databasePassword = "root"
	databaseName     = "database_test"
)

var db *Database

func TestNewDatabase(t *testing.T) {
	config := NewConfig(databaseAddr, databaseUser, databasePassword, databaseName)
	var err error
	if db, err = NewDatabaseWrapper(config, true); err != nil {
		t.Fatal(err)
	}
}

func TestCleanup(t *testing.T) {
	sql := fmt.Sprintf("DROP DATABASE %s", databaseName)
	mdb, err := db.Exec(sql).DB()
	if err != nil {
		t.Fatal(err)
	}
	if err = mdb.Close(); err != nil {
		t.Fatal(err)
	}
}
