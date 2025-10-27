package app

import (
	"errors"
	"fmt"

	"github.com/Pholluxion/task-api/internal/config"
	"github.com/Pholluxion/task-api/internal/model"
	"github.com/Pholluxion/task-api/internal/service"
	"github.com/Pholluxion/task-api/internal/store"
	"github.com/Pholluxion/task-api/internal/transport"
	"github.com/Pholluxion/task-api/internal/transport/middlewares"
	"github.com/Pholluxion/task-api/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type App struct {
	Router *gin.Engine
	Port   string
}

func New() (*App, error) {
	config := config.NewConfig()

	jwtService := utils.NewJWTService(
		config.SecretKey,
		config.TokenExpireTime,
	)

	db, err := gorm.Open(sqlite.Open(config.DBName), &gorm.Config{})

	if err != nil {
		return nil, errors.New("failed to connect database")
	}

	if err := db.AutoMigrate(&model.Task{}); err != nil {
		return nil, err
	}

	taskStore := store.New(db)
	taskService := service.New(&taskStore)
	taskHandler := transport.NewTaskHandler(&taskService)
	authHandler := transport.NewAuthHandler(jwtService)

	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())
	router.GET("/login", authHandler.Login())
	// Rutas públicas
	authorized := router.Group("/api")
	authorized.Use(middlewares.AuthMiddleware(jwtService))
	{
		authorized.GET("/tasks", taskHandler.GetAll())
		authorized.GET("/tasks/:id", taskHandler.GetByID())
		authorized.POST("/tasks", taskHandler.Create())
		authorized.PUT("/tasks/:id", taskHandler.Update())
		authorized.DELETE("/tasks/:id", taskHandler.Delete())
	}

	return &App{Router: router, Port: config.Port}, nil
}

func (a *App) Start() error {
	return a.Router.Run(fmt.Sprintf(":%s", a.Port))
}
