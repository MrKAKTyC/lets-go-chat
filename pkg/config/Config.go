package config

import "os"

type DBConfig struct {
	Username, URL, DBName string
}

type ServerConfig struct {
	Port string
}

type Config struct {
	DB     DBConfig
	Server ServerConfig
}

func InitConfig() Config {
	db := DBConfig{URL: os.Getenv("DATABASE_URL")}
	server := ServerConfig{Port: os.Getenv("PORT")}

	return Config{db, server}
}
