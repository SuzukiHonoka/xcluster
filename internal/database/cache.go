package database

var DB *Database

func InitDatabase(config *Config, create bool) error {
	var err error
	if DB, err = NewDatabase(config, create); err != nil {
		return err
	}
	return nil
}
