package service

import (
	"task-manager-example/internal/domain"
	"task-manager-example/internal/repo"
	"task-manager-example/pkg/logger"

	"go.uber.org/zap"
)

type TaskService struct {
	repo   repo.TaskRepository
	logger *logger.Logger
}

func NewTaskService(repo repo.TaskRepository) *TaskService {
	return &TaskService{
		repo:   repo,
		logger: logger.New(),
	}
}

func (s *TaskService) CreateTask(task domain.Task) (domain.Task, error) {
	s.logger.Debug("Creating new task",
		zap.String("title", task.Title))

	createdTask, err := s.repo.CreateTask(task)
	if err != nil {
		s.logger.Error("Failed to create task in repository",
			zap.String("error", err.Error()),
			zap.String("title", task.Title))
		return domain.Task{}, err
	}

	s.logger.Info("Task created successfully",
		zap.String("task_id", createdTask.ID),
		zap.String("title", createdTask.Title))

	return createdTask, nil
}

func (s *TaskService) GetTask(id string) (domain.Task, error) {
	s.logger.Debug("Getting task",
		zap.String("task_id", id))

	task, err := s.repo.GetTask(id)
	if err != nil {
		s.logger.Error("Failed to get task from repository",
			zap.String("error", err.Error()),
			zap.String("task_id", id))
		return domain.Task{}, err
	}

	s.logger.Debug("Task retrieved successfully",
		zap.String("task_id", task.ID),
		zap.String("title", task.Title))

	return task, nil
}

func (s *TaskService) UpdateTask(task domain.Task) (domain.Task, error) {
	s.logger.Debug("Updating task",
		zap.String("task_id", task.ID),
		zap.String("title", task.Title))

	updatedTask, err := s.repo.UpdateTask(task)
	if err != nil {
		s.logger.Error("Failed to update task in repository",
			zap.String("error", err.Error()),
			zap.String("task_id", task.ID))
		return domain.Task{}, err
	}

	s.logger.Info("Task updated successfully",
		zap.String("task_id", updatedTask.ID),
		zap.String("title", updatedTask.Title))

	return updatedTask, nil
}

func (s *TaskService) DeleteTask(id string) error {
	s.logger.Debug("Deleting task",
		zap.String("task_id", id))

	err := s.repo.DeleteTask(id)
	if err != nil {
		s.logger.Error("Failed to delete task from repository",
			zap.String("error", err.Error()),
			zap.String("task_id", id))
		return err
	}

	s.logger.Info("Task deleted successfully",
		zap.String("task_id", id))

	return nil
}

func (s *TaskService) GetAllTasks() []domain.Task {
	s.logger.Debug("Getting all tasks")

	tasks := s.repo.GetAllTasks()

	s.logger.Debug("Retrieved all tasks",
		zap.Int("count", len(tasks)))

	return tasks
}
