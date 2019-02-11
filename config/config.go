package config

import (
	"github.com/jinzhu/gorm"
)

type Config struct {
	DatabasePath string
	Port         string
}

func NewConfig() *Config {
	return &Config{
		DatabasePath: "/home/joe/go/src/data/ZKDB.db",
		Port:         "8000",
	}
}
func ConnectDatabase(config *Config) (*gorm.DB, error) {
	return gorm.Open("sqlite3", config.DatabasePath)
}
