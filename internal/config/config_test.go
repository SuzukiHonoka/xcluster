package config

import (
	"fmt"
	"math/rand"
	"testing"
	"xcluster/internal/database"
)

var (
	databaseAddr     = "127.0.0.1:3306"
	databaseUser     = "root"
	databasePassword = "root"
	databaseName     = "database_test"
)

var (
	n  = 10
	kv = make(map[string]string, n)
)

func init() {
	// init test database
	config := database.NewConfig(databaseAddr, databaseUser, databasePassword, databaseName)
	if err := database.InitDatabase(config, true); err != nil {
		panic(err)
	}
	if err := InitConfigTable(); err != nil {
		panic(err)
	}
	InitCache()
	//generate kv
	for i := 0; i < 10; i++ {
		kv[randString(20)] = randString(20)
	}
}

func randString(n int) string {
	var table = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890@.")
	s := make([]rune, n)
	for i := range s {
		s[i] = table[rand.Intn(len(table))]
	}
	return string(s)
}

func TestSetConfig(t *testing.T) {
	for k, v := range kv {
		if _, err := SetConfig(k, v); err != nil {
			t.Error(err)
		}
	}
}

func TestKey_GetValue(t *testing.T) {
	for k, v := range kv {
		val, err := KeyName(k).GetValue()
		if err != nil {
			t.Error(err)
		}
		if val != v {
			err = fmt.Errorf("val mismatch, expect=%s, but=%s", v, val)
		}
	}
}

func TestKey_DeleteConfig(t *testing.T) {
	for k := range kv {
		if err := KeyName(k).DeleteConfig(); err != nil {
			t.Error(err)
		}
	}
}

func TestCleanup(t *testing.T) {
	sql := fmt.Sprintf("DROP DATABASE %s", databaseName)
	mdb, err := database.DB.Exec(sql).DB()
	if err != nil {
		t.Fatal(err)
	}
	if err = mdb.Close(); err != nil {
		t.Fatal(err)
	}
}
