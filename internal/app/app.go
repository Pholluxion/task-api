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

	db, err := gorm.Open(sqlite.Open(config.DBName), &gorm.Config{})

	if err != nil {
		return nil, errors.New("failed to connect database")
	}

	if err := db.AutoMigrate(&model.Task{}, &model.User{}); err != nil {
		return nil, err
	}

	taskStore := store.New(db)
	userStore := store.NewUserStore(db)

	taskService := service.New(&taskStore)
	authService := service.NewAuthService(userStore, config.TokenExpireTime, []byte(config.SecretKey))
	taskHandler := transport.NewTaskHandler(&taskService)
	authHandler := transport.NewAuthHandler(authService)

	router := gin.Default()

	router.Use(middlewares.AllowCORS())
	router.GET("/token", authHandler.Login())
	router.POST("/register", authHandler.Register())

	authorized := router.Group("/api")
	authorized.Use(middlewares.JWTAuthMiddleware(authService))
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
