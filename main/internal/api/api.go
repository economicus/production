package api

import (
	"economicus/config"
	"economicus/internal/api/handler"
	"economicus/internal/api/hateos"
	"economicus/internal/api/middleware"
	"economicus/internal/api/repository"
	"economicus/internal/api/routes"
	"economicus/internal/api/service"
	"economicus/internal/api/token"
	"economicus/internal/drivers"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

type Manager struct {
	routes *gin.Engine
	mid    *middleware.AuthMiddleware
	app    *config.AppConfig
	db     *drivers.DB
	aws    *drivers.AWS
	logger *log.Logger
	hateos *hateos.Hateos
	jwt    *token.JwtManager
}

func NewManager(app *config.AppConfig, db *drivers.DB) *Manager {
	jwtConf := config.NewJwtConfig()
	jwtMan := token.NewJwtTokenManager(jwtConf)
	aws := drivers.NewAWS()
	h := hateos.NewHateos(app)
	authMid := middleware.NewAuthMiddleware(db, jwtMan)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	return &Manager{
		routes: r,
		mid:    authMid,
		app:    app,
		db:     db,
		hateos: h,
		aws:    aws,
		logger: NewLogger(),
		jwt:    jwtMan,
	}
}

func (m *Manager) Run(port string) error {
	return m.routes.Run(fmt.Sprintf(":%s", port))
}

func (m *Manager) Setup() {
	m.setupAuth()
	m.setupUser()
	m.setupQuant()
	m.setupComment()
	m.setupReply()
}

func (m *Manager) setupAuth() {
	repo := repository.NewAuthRepository(m.db.SQL, m.jwt, m.logger)
	serv := service.NewAuthService(repo)
	hdr := handler.NewAuthHandler(serv)
	r := m.routes.Group("/v1")
	rt := routes.NewAuthRoute(r, hdr)
	rt.Setup()
}

func (m *Manager) setupUser() {
	repo := repository.NewUserRepository(m.db.SQL, m.aws, m.logger)
	serv := service.NewUserService(repo, m.aws)
	hdr := handler.NewUserHandler(serv, m.hateos)
	r := m.routes.Group("/v1")
	rt := routes.NewUserRoute(r, hdr, m.mid)
	rt.Setup()
}

func (m *Manager) setupQuant() {
	repo := repository.NewQuantRepository(m.db.SQL, m.logger)
	serv := service.NewQuantService(repo)
	hdr := handler.NewQuantHandler(serv)
	r := m.routes.Group("/v1", m.mid.Authenticate())
	rt := routes.NewQuantRoute(r, hdr)
	rt.Setup()
}

func (m *Manager) setupComment() {
	repo := repository.NewCommentRepository(m.db.SQL, m.logger)
	serv := service.NewCommentService(repo)
	hdr := handler.NewCommentHandler(serv, m.hateos)
	r := m.routes.Group("/v1", m.mid.Authenticate())
	rt := routes.NewCommentRoute(r, hdr)
	rt.Setup()
}

func (m *Manager) setupReply() {
	repo := repository.NewReplyRepository(m.db.SQL, m.logger)
	serv := service.NewReplyService(repo)
	hdr := handler.NewReplyHandler(serv, m.hateos)
	r := m.routes.Group("/v1", m.mid.Authenticate())
	rt := routes.NewReplyRoute(r, hdr)
	rt.Setup()
}
