package config

import (
	"os"
)

type Config struct {
	MongoURI string
	Database string
}

func LoadConfig() *Config {
	return &Config{
		MongoURI: getEnv("MONGO_URI", "mongodb://admin:secret@localhost:27017/"),
		Database: getEnv("MONGO_DATABASE", "catalogdb"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
