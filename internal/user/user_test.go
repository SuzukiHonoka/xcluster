package user

import (
	"fmt"
	"math/rand"
	"strconv"
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
	n     = 10
	users []User
)

func init() {
	// init test database
	config := database.NewConfig(databaseAddr, databaseUser, databasePassword, databaseName)
	if err := database.InitDatabase(config, true); err != nil {
		panic(err)
	}
	if err := InitUserTable(); err != nil {
		panic(err)
	}
	//generate users
	users = generateUsers(n)
}

func generateUsers(n int) []User {
	t := make([]User, n)
	for i := range t {
		t[i] = User{
			Name:     Name(randString(100)),
			Password: Password(randString(100)),
			Email:    randString(100),
		}
	}
	return t
}

func randString(n int) string {
	var table = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890@.")
	s := make([]rune, n)
	for i := range s {
		s[i] = table[rand.Intn(len(table))]
	}
	return string(s)
}

func TestNewUser(t *testing.T) {
	for _, user := range users {
		//t.Logf("ID=%d, name=%s, password=%s, email=%s", user.ID, user.Name, user.Password, user.Email)
		_, err := NewUser(string(user.Name), string(user.Password), user.Email)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestGetAll(t *testing.T) {
	all, err := GetAll()
	if err != nil {
		t.Fatal(err)
	}
	//for _, user := range all {
	//	t.Logf("ID=%d, name=%s, password=%s, email=%s", user.ID, user.Name, user.Password, user.Email)
	//}
	if len(all) < n {
		t.Fatalf("data loses, test user count=%d, actuall count=%d", n, len(all))
	}
}

func TestID_GetUser(t *testing.T) {
	all, err := GetAll()
	if err != nil {
		t.Fatal(err)
	}
	for _, user := range all {
		var u *User
		u, err = user.ID.GetUser()
		if err != nil {
			t.Fatal(err)
		}
		if *u != *user {
			t.Fatalf("integrity check failed, original user=%+v, actuall user=%+v", user, u)
		}
	}
}

func TestName_GetUser(t *testing.T) {
	all, err := GetAll()
	if err != nil {
		t.Fatal(err)
	}
	for _, user := range all {
		var u *User
		u, err = user.Name.GetUser()
		if err != nil {
			t.Error(err)
		}
		if *u != *user {
			t.Fatalf("integrity check failed, original user=%+v, actuall user=%+v", user, u)
		}
	}
}

func TestUser_Update(t *testing.T) {
	all, err := GetAll()
	if err != nil {
		t.Fatal(err)
	}
	// set all field to val
	for _, user := range all {
		val := strconv.Itoa(int(user.ID))
		p := &val
		if err = user.Update(p, nil, nil); err != nil {
			t.Fatal(err)
		}
	}
	// verify updates
	all, err = GetAll()
	if err != nil {
		t.Fatal(err)
	}
	for _, user := range all {
		val := strconv.FormatInt(int64(user.ID), 10)
		if string(user.Name) != val {
			t.Fatal("update failed")
		}
	}
}

func TestID_DeleteUser(t *testing.T) {
	all, err := GetAll()
	if err != nil {
		t.Fatal(err)
	}
	for _, user := range all {
		var u *User
		u, err = user.ID.GetUser()
		if err != nil {
			t.Fatal(err)
		}
		if *u != *user {
			t.Fatalf("integrity check failed, original user=%+v, actuall user=%+v", user, u)
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
