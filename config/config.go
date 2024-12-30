package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort       int
	DatabaseHost     string
	DatabasePort     int
	DatabaseUser     string
	DatabasePassword string
	DatabaseDbName   string
}

var appConfig *Config

func defaultIfEmpty(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

func parseInt(s string) int {
	value, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println("Parse int failed:", err.Error())
	}

	return value
}

func init() {
	err := godotenv.Load()
	if err != nil {
		if err.Error() != "open .env: no such file or directory" {
			fmt.Println("Load .env failed:", err.Error())
			os.Exit(1)
		}
	}

	fmt.Println()
	appConfig = &Config{
		ServerPort:       parseInt(defaultIfEmpty(os.Getenv("SERVER_PORT"), "8080")),
		DatabaseHost:     defaultIfEmpty(os.Getenv("DATABASE_HOST"), "localhost"),
		DatabasePort:     parseInt(defaultIfEmpty(os.Getenv("DATABASE_PORT"), "5432")),
		DatabaseUser:     defaultIfEmpty(os.Getenv("DATABASE_USER"), "user"),
		DatabasePassword: defaultIfEmpty(os.Getenv("DATABASE_PASSWORD"), "pass"),
		DatabaseDbName:   defaultIfEmpty(os.Getenv("DATABASE_DB_NAME"), "meme-coin"),
	}
}

func GetConfig() *Config {
	return appConfig
}
