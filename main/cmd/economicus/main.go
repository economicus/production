package main

import (
	"economicus/config"
	"economicus/internal/api"
	"economicus/internal/drivers"
	"log"
)

const port = "8080"

type App struct {
	config   *config.AppConfig
	database *drivers.DB
	routes   *api.Manager
}

func NewApp() *App {
	c := config.NewAppConfig()
	db := drivers.NewDatabase()
	routes := api.NewManager(c, db)

	return &App{
		config:   c,
		database: db,
		routes:   routes,
	}
}

func (app *App) Setup() {
	app.routes.Setup()
}

func (app *App) Run() {
	if err := app.routes.Run(port); err != nil {
		log.Fatalf("error while running server on %s", port)
	}
}

func (app *App) PrintInfo() {
	app.config.PrintInfo()
	app.database.Config.PrintInfo()
}

func main() {
	app := NewApp()
	if !app.config.InProduction {
		app.PrintInfo()
	}
	app.Setup()
	app.Run()
}
