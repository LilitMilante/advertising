package app

import (
	"database/sql"
	"fmt"

	"github.com/LilitMilante/advertising/internal/dal"
)

type Config struct {
	HTTPPort string

	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

func NewConfig() *Config {
	return &Config{
		HTTPPort:   "8081",
		DBHost:     "localhost",
		DBPort:     8080,
		DBUser:     "postgres",
		DBPassword: "dev",
		DBName:     "postgres",
	}
}

type App struct {
	c Config
}

func NewApp(c Config) (*App, error) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	_ = dal.NewRepository(db)

	//	todo Init other modules.

	return &App{}, nil
}

func ConnectDB(c Config) (*sql.DB, error) {
	//todo Make this func as a method of App.
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
