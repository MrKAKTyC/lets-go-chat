package config

import "os"

type dbConfig struct {
	Username, URL, DBName string
}

type serverConfig struct {
	Port string
}

type Config struct {
	DB     dbConfig
	Server serverConfig
}

func InitConfig() Config {
	db := dbConfig{URL: os.Getenv("DATABASE_URL")}
	server := serverConfig{Port: os.Getenv("PORT")}

	return Config{db, server}
}

func LocalConfig() Config {
	db := dbConfig{URL: "postgres://postgres:admin@127.0.0.1:5432/letsGoChat?sslmode=disable"}
	server := serverConfig{Port: "80"}

	return Config{db, server}
}
