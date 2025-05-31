package memory

import (
	"sync"
	"task-manager-example/internal/domain"
	"task-manager-example/internal/repo"
	"time"

	"github.com/google/uuid"
)

type TaskRepository struct {
	tasks map[string]domain.Task
	mu    sync.RWMutex // READ AND WRITE LOCK
}

func NewTaskRepository() repo.TaskRepository {
	return &TaskRepository{
		tasks: make(map[string]domain.Task),
	}
}

func (r *TaskRepository) CreateTask(task domain.Task) (domain.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	task.ID = uuid.New().String()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	r.tasks[task.ID] = task
	return task, nil
}

func (r *TaskRepository) GetTask(id string) (domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	task, ok := r.tasks[id]

	if !ok {
		return domain.Task{}, domain.ErrTaskNotFound
	}
	return task, nil
}

func (r *TaskRepository) DeleteTask(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.tasks[id]; !ok {
		return domain.ErrTaskNotFound
	}
	delete(r.tasks, id)
	return nil
}

func (r *TaskRepository) UpdateTask(task domain.Task) (domain.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.tasks[task.ID]; !ok {
		return domain.Task{}, domain.ErrTaskNotFound
	}
	task.UpdatedAt = time.Now()
	r.tasks[task.ID] = task
	return task, nil
}

func (r *TaskRepository) GetAllTasks() []domain.Task {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]domain.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}
