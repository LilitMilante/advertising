package app

import (
	"database/sql"
	"fmt"

	"github.com/LilitMilante/advertising/internal/api"
	"github.com/LilitMilante/advertising/internal/dal"
	"github.com/LilitMilante/advertising/internal/services"

	_ "github.com/lib/pq"
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
	c      Config
	server *api.Server
}

func NewApp(c Config) (*App, error) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	repo := dal.NewRepository(db)
	service := services.NewAnnouncement(repo)
	handler := api.NewHandler(service)
	server := api.NewServer(c.HTTPPort, handler)

	return &App{
		c:      Config{},
		server: server,
	}, nil
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

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (a *App) Start() error {
	return a.server.ListenAndServe()
}
