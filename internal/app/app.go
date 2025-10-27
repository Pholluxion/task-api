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
	"github.com/Pholluxion/task-api/internal/transport/router"
	"github.com/Pholluxion/task-api/internal/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type App struct {
	Server *http.Server
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

	r := router.New()

	r.Use(middlewares.LoggingMiddleware, middlewares.CORSMiddleware)

	// Rutas públicas
	auth := r.Group("/auth")
	auth.Get("/login", authHandler.Login)

	// Rutas privadas
	api := r.Group("/api")
	api.Use(middlewares.AuthMiddleware(jwtService))

	api.Get("/tasks", taskHandler.GetAll)
	api.Get("/tasks/{id}", taskHandler.GetByID)
	api.Post("/tasks", taskHandler.Create)
	api.Put("/tasks/{id}", taskHandler.Update)
	api.Delete("/tasks/{id}", taskHandler.Delete)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Port),
		Handler: r.Handler(),
	}
	return &App{Server: server}, nil
}

func (a *App) Start() error {
	fmt.Println("✅ Server started on http://localhost:8080")
	return a.Server.ListenAndServe()
}
