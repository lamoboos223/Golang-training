package service

import (
	"context"
	"task-manager-example/internal/domain"
	"task-manager-example/internal/repo"
	"task-manager-example/pkg/logger"
	"time"

	"go.uber.org/zap"
)

type TaskProcessor struct {
	repo   repo.TaskRepository
	logger *logger.Logger
	tasks  chan domain.Task
	done   chan struct{}
}

func NewTaskProcessor(repo repo.TaskRepository) *TaskProcessor {
	return &TaskProcessor{
		repo:   repo,
		logger: logger.New(),
		tasks:  make(chan domain.Task, 100),
		done:   make(chan struct{}),
	}
}

func (p *TaskProcessor) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case task := <-p.tasks:
				time.Sleep(2000 * time.Millisecond)
				task.Status = true
				_, err := p.repo.UpdateTask(task)
				if err != nil {
					p.logger.Error("Failed to process task", zap.Error(err))
					continue
				}
				p.logger.Info("Task processed", zap.String("id", task.ID), zap.Bool("status", task.Status))
			case <-ctx.Done():
				close(p.done)
				return
			}
		}
	}()
}

func (p *TaskProcessor) ProcessTask(task domain.Task) {
	p.tasks <- task
}
