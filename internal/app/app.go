package app

import (
	"errors"
	"fmt"
	"net/http"

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
	Server *http.Server
}

func New() (*App, error) {
	config := config.NewConfig()
	jwtUtil := utils.NewJWTUtil(config.SecretKey, config.TokenExpireTime)

	db, err := gorm.Open(sqlite.Open(config.DBName), &gorm.Config{})

	if err != nil {
		return nil, errors.New("failed to connect database")
	}

	if err := db.AutoMigrate(&model.Task{}, &model.User{}); err != nil {
		return nil, err
	}

	taskStore := store.NewTaskStore(db)
	userStore := store.NewUserStore(db)

	taskService := service.NewTaskService(&taskStore)
	authService := service.NewAuthService(userStore)
	taskHandler := transport.NewTaskHandler(&taskService)
	authHandler := transport.NewAuthHandler(authService, jwtUtil)

	router := gin.Default()

	router.Use(middlewares.AllowCORS())
	router.GET("/token", authHandler.Login())
	router.POST("/register", authHandler.Register())

	authorized := router.Group("/api")
	authorized.Use(middlewares.JWTAuthMiddleware(jwtUtil))
	{
		authorized.GET("/tasks", taskHandler.GetAll())
		authorized.GET("/tasks/:id", taskHandler.GetByID())
		authorized.POST("/tasks", taskHandler.Create())
		authorized.PUT("/tasks/:id", taskHandler.Update())
		authorized.DELETE("/tasks/:id", taskHandler.Delete())
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Port),
		Handler: router.Handler(),
	}

	return &App{Server: server}, nil
}
