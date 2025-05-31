package repo

import "task-manager-example/internal/domain"

type TaskRepository interface {
	CreateTask(task domain.Task) (domain.Task, error)
	GetTask(id string) (domain.Task, error)
	UpdateTask(task domain.Task) (domain.Task, error)
	DeleteTask(id string) error
	GetAllTasks() []domain.Task
}
