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
		Port:  GetEnvStr("PORT", "3000"),
		DBUrl: GetEnvStr("DB_URL", ""),
	}
}

func GetEnvStr(key string, defaultVal string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultVal
	}
	return value
}
