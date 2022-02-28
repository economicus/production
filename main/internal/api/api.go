package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"main/internal/api/middleware"
	"main/internal/conf"
	"main/internal/conf/db/mysql"
	"time"
)

var authMid *middleware.AuthMiddleware

type Router struct {
	engine *gin.Engine
	app    *conf.App
	db     *mysql.DB
}

func New(app *conf.App, db *mysql.DB) *Router {
	e := getEngine()
	authMid = middleware.NewAuthMiddleware(db)
	r := Router{
		engine: e,
		app:    app,
		db:     db,
	}
	r.setAll()
	return &r
}

func (r *Router) Run() {
	if err := r.engine.Run(":" + r.app.InsecurePort); err != nil {
		log.Panicf("error while running app: %v", err)
	}
}

func (r *Router) getGroup() *gin.RouterGroup {
	return r.engine.Group("/v1")
}

func (r *Router) getGroupWithAuth() *gin.RouterGroup {
	return r.engine.Group("/v1", authMid.Authenticate())
}

func getEngine() *gin.Engine {
	e := gin.Default()
	e.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://www.economicus.kr"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	return e
}
