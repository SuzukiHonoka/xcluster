package config

import (
	"errors"
	"gorm.io/gorm"
	"xcluster/internal/database"
)

type Config struct {
	ID      int     `gorm:"type:uint;primaryKey;autoIncrement;unique"`
	KeyName KeyName `gorm:"type:varchar(100);unique;not null"`
	Value   string  `gorm:"type:varchar(100);not null"`
}

func SetConfig(key string, value string) (*Config, error) {
	cache.Set(key, value)
	config := &Config{
		KeyName: KeyName(key),
		Value:   value,
	}
	// check if key already exist
	if err := database.DB.Model(&Config{}).Where("key_name = ?", key).Update("value", value).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err = database.DB.Create(config).Error; err != nil {
				return nil, err
			}
		}
		return nil, err
	}
	return config, nil
}

func (c *Config) Delete() error {
	return c.KeyName.DeleteConfig()
}
