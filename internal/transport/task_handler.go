package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Pholluxion/task-api/internal/model"
	"github.com/Pholluxion/task-api/internal/service"
)

type TaskHandler struct {
	service service.TaskService
}

func New(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: *service}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(15*time.Second))
	defer cancel()
	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdTask, err := h.service.Create(ctx, &task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTask)

}

func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(15*time.Second))
	defer cancel()
	tasks, err := h.service.GetAll(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	statusOK(w, tasks)
}

func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(15*time.Second))
	defer cancel()
	rawID := r.PathValue("id")
	ID, err := strconv.Atoi(rawID)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	task, err := h.service.GetByID(ctx, uint(ID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	statusOK(w, task)

}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(15*time.Second))
	defer cancel()
	rawID := r.PathValue("id")
	ID, err := strconv.Atoi(rawID)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t, err := h.service.Update(ctx, uint(ID), task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	statusOK(w, t)
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(15*time.Second))
	defer cancel()
	rawID := r.PathValue("id")
	ID, err := strconv.Atoi(rawID)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(ctx, uint(ID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func statusOK(w http.ResponseWriter, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
