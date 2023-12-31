package user

import (
	"fmt"
	"strconv"
	"testing"
	"xcluster/internal/database"
	"xcluster/pkg/random"
)

var (
	databaseAddr     = "127.0.0.1:3306"
	databaseUser     = "root"
	databasePassword = "root"
	databaseName     = "database_test"
)

var (
	n     = 10
	users Users
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

func generateUsers(n int) Users {
	t := make(Users, n)
	for i := range t {
		t[i] = &User{
			Name:     Name(random.String(100)),
			Password: Password(random.String(100)),
			Email:    Email(random.String(100)),
		}
	}
	return t
}

func TestNewUser(t *testing.T) {
	for _, user := range users {
		//t.Logf("ID=%d, name=%s, password=%s, email=%s", user.ID, user.Name, user.Password, user.Email)
		_, err := NewUser(string(user.Name), string(user.Password), string(user.Email))
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestGetAll(t *testing.T) {
	all, err := All()
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
	all, err := All()
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
	all, err := All()
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
	all, err := All()
	if err != nil {
		t.Fatal(err)
	}
	// set all field to val
	for _, user := range all {
		val := strconv.Itoa(int(user.ID))
		if err = user.Update(val, "", ""); err != nil {
			t.Fatal(err)
		}
	}
	// verify updates
	all, err = All()
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
	all, err := All()
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
