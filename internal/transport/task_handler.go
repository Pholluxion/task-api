package transport

import (
	"context"
	"net/http"
	"time"

	"github.com/Pholluxion/task-api/internal/model"
	"github.com/Pholluxion/task-api/internal/service"
	"github.com/Pholluxion/task-api/internal/transport/httpx"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: *service}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(15*time.Second))
	defer cancel()
	var task model.Task
	if err := httpx.Decode(r, &task); err != nil {
		httpx.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	createdTask, err := h.service.Create(ctx, &task)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpx.JSON(w, http.StatusCreated, createdTask)

}

func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(15*time.Second))
	defer cancel()
	tasks, err := h.service.GetAll(ctx)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpx.JSON(w, http.StatusOK, tasks)
}

func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(15*time.Second))
	defer cancel()

	id, err := httpx.ParamID(r)

	if err != nil {
		httpx.Error(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	task, err := h.service.GetByID(ctx, id)

	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpx.JSON(w, http.StatusOK, task)
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(15*time.Second))
	defer cancel()

	id, err := httpx.ParamID(r)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var task model.Task
	if err := httpx.Decode(r, &task); err != nil {
		httpx.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	t, err := h.service.Update(ctx, id, task)

	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpx.JSON(w, http.StatusOK, t)
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(15*time.Second))
	defer cancel()
	id, err := httpx.ParamID(r)
	if err != nil {
		httpx.Error(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.service.Delete(ctx, id); err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpx.JSON(w, http.StatusOK, map[string]string{"message": "Task deleted successfully"})
}
