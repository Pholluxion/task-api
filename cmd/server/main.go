package main

import (
	"fmt"
	"net/http"

	"github.com/Pholluxion/task-api/internal/model"
	"github.com/Pholluxion/task-api/internal/service"
	"github.com/Pholluxion/task-api/internal/store"
	"github.com/Pholluxion/task-api/internal/transport"
	"github.com/Pholluxion/task-api/internal/transport/middlewares"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DBName = "tasks.db"

func main() {

	db, err := gorm.Open(sqlite.Open(DBName), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Task{})

	taskStore := store.New(db)
	taskService := service.New(&taskStore)
	taskHandler := transport.New(&taskService)
	authHandler := transport.AuthHandler{}

	mux := http.NewServeMux()

	mux.Handle("GET /tasks", middlewares.AuthMiddleware(http.HandlerFunc(taskHandler.GetAll)))
	mux.Handle("GET /tasks/{id}", middlewares.AuthMiddleware(http.HandlerFunc(taskHandler.GetByID)))
	mux.Handle("POST /tasks", middlewares.AuthMiddleware(http.HandlerFunc(taskHandler.Create)))
	mux.Handle("PUT /tasks/{id}", middlewares.AuthMiddleware(http.HandlerFunc(taskHandler.Update)))
	mux.Handle("DELETE /tasks/{id}", middlewares.AuthMiddleware(http.HandlerFunc(taskHandler.Delete)))
	mux.HandleFunc("/login", authHandler.Login)

	server := &http.Server{
		Addr:    ":8080",
		Handler: middlewares.CORSMiddleware(middlewares.LoggingMiddleware(mux)),
	}

	fmt.Println("✅ Server started on http://localhost:8080")

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("❌ Failed to start server:", err)
	}
}
