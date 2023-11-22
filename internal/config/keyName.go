package config

import "xcluster/internal/database"

type KeyName string

const (
	KeyNameInstalled KeyName = "Installed"
	KeyNameURL       KeyName = "URL"
	KeyNameTitle
)

func (k KeyName) GetConfig() (*Config, error) {
	var config Config
	if err := database.DB.First(&config, "key_name = ?", k).Error; err != nil {
		return nil, err
	}
	return &config, nil
}

func (k KeyName) GetValue() (string, error) {
	// get from cache first
	if val, err := cache.Get(string(k)); err == nil {
		return val, nil
	}
	config, err := k.GetConfig()
	if err != nil {
		return "", err
	}
	cache.Set(string(k), config.Value)
	return config.Value, nil
}

func (k KeyName) DeleteConfig() error {
	return database.DB.Delete(&Config{}, "key_name = ?", k).Error
}
