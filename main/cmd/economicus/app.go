package main

import (
	"main/internal/api"
	"main/internal/conf"
	db "main/internal/conf/db/mysql"
)

type App struct {
	config *conf.App
	mysql  *db.MySql
	router *api.Router
}

func New() *App {
	c := conf.New()
	mysql := db.NewMySql()
	router := api.New(c, mysql)

	return &App{
		config: c,
		mysql:  mysql,
		router: router,
	}
}

func (app *App) Run() {
	app.router.Run()
}
