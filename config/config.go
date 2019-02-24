package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	DatabasePath       string
	SystemDatabasePath string
	Port               string
	DeviceSerial       string
	EnableLog          bool
	Logfile            string
	MaxLogDays         int
}

func NewConfig() *Config {
	return &Config{
		DatabasePath:       "/home/joe/go/src/data/ZKDB.db",
		SystemDatabasePath: "/home/joe/go/src/data/ZKSystem.db",
		Port:               "8000",
	}
}
func ConnectDatabase(config Config) (*gorm.DB, error) {
	return gorm.Open("sqlite3", config.DatabasePath)
}
