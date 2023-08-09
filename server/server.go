package server

import (
	"log"

	"github.com/k2sebeom/go-echo-server-template/api/routes"
	"github.com/k2sebeom/go-echo-server-template/config"
	_ "github.com/k2sebeom/go-echo-server-template/docs"
	"github.com/k2sebeom/go-echo-server-template/middlewares"
	"github.com/k2sebeom/go-echo-server-template/repositories"
	"github.com/k2sebeom/go-echo-server-template/services"
	"github.com/k2sebeom/go-echo-server-template/utils"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//	@title			Ecgi Server API
//	@version		1.0
//	@description	Echo Backend Server

//	@securityDefinitions.apikey	JWT
//	@in							header
//	@name						Authorization
type Server struct {
	*echo.Echo
	*config.Config
}

func NewServer(

) (*Server, error) {
	e := echo.New()
	c := config.GlobalConfig
	if c == nil {
		log.Fatal("Failed to load configs")
	}
	// Initialize Middlewares
	e.Use(middlewares.Logger())
	e.Use(middlewares.NewCorsMiddleware())

	// API DOC
	e.GET("/docs/*", echoSwagger.WrapHandler)

	// Assign Validator
	v := utils.NewRequestValidator()
	e.Validator = v

	// Initialize DB
	db, err := utils.NewDBInstance(*c);
	
	if err != nil {
		return nil, err
	}
	// Initialize Repositories
	userRepo := repositories.NewUserRepoImpl(db)
	authRepo := repositories.NewAuthRepositoryImpl(db)

	// Services
	healthService := services.NewHealthService()
	authService := services.NewAuthService(userRepo, authRepo)
	oauthService := services.NewOAuthService(userRepo)

	// Initialize Routes
	router := e.Group("/api")
	routes.HealthRoutes(router, healthService)
	routes.AuthRoutes(router, authService, oauthService)

	return &Server{
		Echo: e, Config: c,
	}, nil
}
