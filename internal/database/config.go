package database

import "fmt"

const (
	// use tcp as default protocol
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for dns format details
	dsnFormat       = "%s:%s@tcp(%s)/"
	dsnSelectFormat = dsnFormat + "%s?charset=utf8mb4&parseTime=True&loc=Local"
)

type Config struct {
	Addr         string // host:port
	User         string
	Password     string
	DatabaseName string // database name
}

func NewConfig(addr, user, password, name string) *Config {
	return &Config{
		Addr:         addr,
		User:         user,
		Password:     password,
		DatabaseName: name,
	}
}

func (c *Config) RootDSN() string {
	return c.generateDSN(false)
}

func (c *Config) DSN() string {
	return c.generateDSN(true)
}

func (c *Config) generateDSN(selectDatabase bool) string {
	var dsn string
	if selectDatabase {
		dsn = fmt.Sprintf(dsnSelectFormat, c.User, c.Password, c.Addr, c.DatabaseName)
	} else {
		dsn = fmt.Sprintf(dsnFormat, c.User, c.Password, c.Addr)
	}
	return dsn
}
