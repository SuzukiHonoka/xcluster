package database

import (
	"fmt"
	"testing"
)

var (
	addr     = "127.0.0.1:3306"
	user     = "root"
	password = "root"
	name     = "database_test"
)

func TestNewDatabase(t *testing.T) {
	config := NewConfig(addr, user, password, name)
	db, err := NewDatabase(config, true)
	if err != nil {
		t.Fatal(err)
	}
	sql := fmt.Sprintf("DROP DATABASE %s", config.DatabaseName)
	mdb, err := db.Exec(sql).DB()
	if err != nil {
		t.Fatal(err)
	}
	if err = mdb.Close(); err != nil {
		t.Fatal(err)
	}
}
