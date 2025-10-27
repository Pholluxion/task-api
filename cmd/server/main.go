package main

import (
	"fmt"
	"net/http"

	"task-api/internal/db"
	"task-api/internal/service"
	"task-api/internal/store"
	"task-api/internal/transport"
	"task-api/internal/transport/middlewares"
)

func main() {

	db := db.New()
	taskStore := store.New(db.DB)
	taskService := service.New(&taskStore)
	taskHandler := transport.New(&taskService)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /tasks", taskHandler.GetAll)
	mux.HandleFunc("GET /tasks/{id}", taskHandler.GetByID)
	mux.HandleFunc("POST /tasks", taskHandler.Create)
	mux.HandleFunc("PUT /tasks/{id}", taskHandler.Update)
	mux.HandleFunc("DELETE /tasks/{id}", taskHandler.Delete)

	server := &http.Server{
		Addr:    ":8080",
		Handler: middlewares.CORSMiddleware(middlewares.LoggingMiddleware(mux)),
	}

	fmt.Println("✅ Server started on http://localhost:8080")

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("❌ Failed to start server:", err)
	}
}
