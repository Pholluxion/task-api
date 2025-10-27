package transport

import (
	"net/http"

	"github.com/Pholluxion/task-api/internal/model"
	"github.com/Pholluxion/task-api/internal/service"
	"github.com/Pholluxion/task-api/internal/transport/httpx"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: *service}
}

func (h *TaskHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var task model.Task
		if err := ctx.ShouldBindJSON(&task); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		createdTask, err := h.service.Create(ctx, &task)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, createdTask)
	}

}

func (h *TaskHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tasks, err := h.service.GetAll(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, tasks)
	}
}

func (h *TaskHandler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := httpx.ParamID(ctx.Request)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		task, err := h.service.GetByID(ctx, id)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, task)
	}
}

func (h *TaskHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := httpx.ParamID(ctx.Request)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var task model.Task

		if err := ctx.ShouldBindJSON(&task); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		t, err := h.service.Update(ctx, id, task)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, t)
	}

}

func (h *TaskHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := httpx.ParamID(ctx.Request)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := h.service.Delete(ctx, id); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, map[string]string{"message": "Task deleted successfully"})
	}
}
