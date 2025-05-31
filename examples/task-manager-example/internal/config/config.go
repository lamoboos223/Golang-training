package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	RepositoryType string `yaml:"repository_type"` // "memory" or "database"
	Database       DatabaseConfig
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

func NewConfig() *Config {
	// Default config
	cfg := &Config{
		RepositoryType: "memory",
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "postgres",
			DBName:   "taskmanager",
		},
	}

	// Try to load from config file
	if file, err := os.Open("config.yaml"); err == nil {
		defer file.Close()
		yaml.NewDecoder(file).Decode(cfg)
	}

	return cfg
}
