package config

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

type Config struct {
	Port  string
	DBUrl string
}

func LoadConfig() Config {
	return Config{
		Port:  GetEnvStr("PORT", "8080"),
		DBUrl: GetEnvStr("DATABASE_URL", ""),
	}
}

func GetEnvStr(key string, defaultVal string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultVal
	}
	return value
}
