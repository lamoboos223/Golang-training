package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"task-manager-example/internal/domain"
	"task-manager-example/internal/service"
	"task-manager-example/pkg/logger"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type TaskHandler struct {
	service *service.TaskService
	logger  *logger.Logger
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
		logger:  logger.New(),
	}
}

// AddRequestID adds a unique request ID to the context
func AddRequestID(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		ctx := context.WithValue(r.Context(), "request_id", requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	// Get logger with request context
	log := h.logger.WithContext(r.Context())

	var task domain.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		log.Error("Failed to decode request body",
			zap.String("error", err.Error()))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdTask, err := h.service.CreateTask(task)
	if err != nil {
		log.Error("Failed to create task",
			zap.String("error", err.Error()),
			zap.String("title", task.Title))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Info("Task created successfully",
		zap.String("task_id", createdTask.ID),
		zap.String("title", createdTask.Title))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTask)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		h.logger.Warn("Missing task ID in request")
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	task, err := h.service.GetTask(id)
	if err != nil {
		if err == domain.ErrTaskNotFound {
			h.logger.Warn("Task not found",
				zap.String("task_id", id))
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		h.logger.Error("Failed to get task",
			zap.String("error", err.Error()),
			zap.String("task_id", id))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.logger.Info("Task retrieved successfully",
		zap.String("task_id", task.ID),
		zap.String("title", task.Title))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task domain.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		h.logger.Error("Failed to decode request body",
			zap.String("error", err.Error()))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedTask, err := h.service.UpdateTask(task)
	if err != nil {
		if err == domain.ErrTaskNotFound {
			h.logger.Warn("Task not found for update",
				zap.String("task_id", task.ID))
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		h.logger.Error("Failed to update task",
			zap.String("error", err.Error()),
			zap.String("task_id", task.ID))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.logger.Info("Task updated successfully",
		zap.String("task_id", updatedTask.ID),
		zap.String("title", updatedTask.Title))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		h.logger.Warn("Missing task ID in request")
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteTask(id)
	if err != nil {
		if err == domain.ErrTaskNotFound {
			h.logger.Warn("Task not found for deletion",
				zap.String("task_id", id))
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		h.logger.Error("Failed to delete task",
			zap.String("error", err.Error()),
			zap.String("task_id", id))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.logger.Info("Task deleted successfully",
		zap.String("task_id", id))

	w.WriteHeader(http.StatusNoContent)
}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.service.GetAllTasks()

	h.logger.Info("Retrieved all tasks",
		zap.Int("count", len(tasks)))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
