package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	AppEnv    string
	AppPort   string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	DBSSL     string
	JwtSecret string
}

func NewConfigFromEnv() (*Config, error) {
	cfg := &Config{
		AppEnv:    os.Getenv("APP_ENV"),
		AppPort:   os.Getenv("APP_PORT"),
		DBHost:    os.Getenv("DB_HOST"),
		DBPort:    os.Getenv("DB_PORT"),
		DBUser:    os.Getenv("DB_USER"),
		DBPass:    os.Getenv("DB_PASSWORD"),
		DBName:    os.Getenv("DB_NAME"),
		DBSSL:     os.Getenv("DB_SSLMODE"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}

	// Validasi: semua variabel wajib ada
	if cfg.AppEnv == "" {
		return nil, errors.New("APP_ENV is required")
	}
	if cfg.AppPort == "" {
		return nil, errors.New("APP_PORT is required")
	}
	if cfg.DBHost == "" {
		return nil, errors.New("DB_HOST is required")
	}
	if cfg.DBPort == "" {
		return nil, errors.New("DB_PORT is required")
	}
	if cfg.DBUser == "" {
		return nil, errors.New("DB_USER is required")
	}
	if cfg.DBPass == "" {
		return nil, errors.New("DB_PASSWORD is required")
	}
	if cfg.DBName == "" {
		return nil, errors.New("DB_NAME is required")
	}
	if cfg.DBSSL == "" {
		return nil, errors.New("DB_SSLMODE is required")
	}
	if cfg.JwtSecret == "" {
		return nil, errors.New("JWT_SECRET is required")
	}

	return cfg, nil
}

// PostgresDSN menghasilkan DSN untuk koneksi ke PostgreSQL
func (c *Config) PostgresDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		c.DBHost, c.DBUser, c.DBPass, c.DBName, c.DBPort, c.DBSSL,
	)
}

// AppPortInt mengembalikan APP_PORT dalam bentuk integer
func (c *Config) AppPortInt() (int, error) {
	return strconv.Atoi(c.AppPort)
}
