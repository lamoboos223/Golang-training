package factory

import (
	"task-manager-example/internal/config"
	"task-manager-example/internal/repo"
	"task-manager-example/internal/repo/database"
	"task-manager-example/internal/repo/memory"
)

func NewRepository(cfg *config.Config) (repo.TaskRepository, error) {
	switch cfg.RepositoryType {
	case "postgres":
		return database.NewTaskRepository(cfg)
	default:
		return memory.NewTaskRepository(), nil
	}
}
