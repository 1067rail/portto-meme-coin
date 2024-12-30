package config

type Config struct {
	ServerPort       int
	DatabaseHost     string
	DatabasePort     int
	DatabaseUser     string
	DatabasePassword string
	DatabaseDbName   string
}

var appConfig *Config

func init() {
	appConfig = &Config{
		ServerPort:       8080,
		DatabaseHost:     "localhost",
		DatabasePort:     5432,
		DatabaseUser:     "user",
		DatabasePassword: "pass",
		DatabaseDbName:   "meme-coin",
	}
}

func GetConfig() *Config {
	return appConfig
}
