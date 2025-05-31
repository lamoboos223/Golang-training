package api

import (
	"net/http"
	"task-manager-example/internal/api/handlers"
)

func SetupRouter(taskHandler *handlers.TaskHandler) http.Handler {
	mux := http.NewServeMux()

	// Add request ID middleware to all routes
	mux.HandleFunc("/tasks", handlers.AddRequestID(taskHandler.CreateTask))
	mux.HandleFunc("/task", handlers.AddRequestID(taskHandler.GetTask))
	mux.HandleFunc("/task/update", handlers.AddRequestID(taskHandler.UpdateTask))
	mux.HandleFunc("/task/delete", handlers.AddRequestID(taskHandler.DeleteTask))
	mux.HandleFunc("/tasks/all", handlers.AddRequestID(taskHandler.GetAllTasks))

	return mux
}
