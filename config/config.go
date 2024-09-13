package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress    string
	PostgresConn     string
	PostgresJDBCURL  string
	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     string
	PostgresDB       string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Reading from environment variables.")
	}

	cfg := &Config{
		ServerAddress:    os.Getenv("SERVER_ADDRESS"),
		PostgresConn:     os.Getenv("POSTGRES_CONN"),
		PostgresJDBCURL:  os.Getenv("POSTGRES_JDBC_URL"),
		PostgresUser:     os.Getenv("POSTGRES_USERNAME"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresDB:       os.Getenv("POSTGRES_DATABASE"),
	}

	if cfg.ServerAddress == "" {
		return nil, fmt.Errorf("SERVER_ADDRESS is required")
	}
	if cfg.PostgresConn == "" && (cfg.PostgresUser == "" || cfg.PostgresPassword == "" || cfg.PostgresHost == "" || cfg.PostgresPort == "" || cfg.PostgresDB == "") {
		return nil, fmt.Errorf("PostgreSQL connection details are missing")
	}

	return cfg, nil
}
