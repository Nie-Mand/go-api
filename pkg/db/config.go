package db

type DBConfig struct {
	URL string
}

func NewDBConfig(url string) DBConfig {
	return DBConfig{
		URL: url,
	}
}

func NewDefaultDBConfig() DBConfig {
	return NewDBConfig("")
}
