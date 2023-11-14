package redis

type Config struct {
	Addr     string
	Password string
}

func NewConfig(addr, password string) *Config {
	return &Config{
		Addr:     addr,
		Password: password,
	}
}
