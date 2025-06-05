package database

import (
	"database/sql"
	"fmt"
	"sync"
	"task-manager-example/internal/config"
	"task-manager-example/internal/domain"
	"task-manager-example/internal/repo"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// this should be the implementation of the TaskRepository interface
type TaskRepository struct {
	db *sql.DB
	mu sync.RWMutex
}

func NewTaskRepository(cfg *config.Config) (repo.TaskRepository, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Create tasks table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id VARCHAR(255) PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description TEXT,
			status BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		)
	`)
	if err != nil {
		return nil, err
	}

	return &TaskRepository{db: db, mu: sync.RWMutex{}}, nil
}

func (r *TaskRepository) CreateTask(task domain.Task) (domain.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	// Generate new ID and timestamps
	task.ID = uuid.New().String()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	query := `
		INSERT INTO tasks (id, title, description, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, title, description, status, created_at, updated_at
	`

	err := r.db.QueryRow(
		query,
		task.ID,
		task.Title,
		task.Description,
		task.Status,
		task.CreatedAt,
		task.UpdatedAt,
	).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	return task, err
}

func (r *TaskRepository) GetTask(id string) (domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var task domain.Task
	query := `SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return domain.Task{}, domain.ErrTaskNotFound
	}

	return task, err
}

func (r *TaskRepository) UpdateTask(task domain.Task) (domain.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	query := `
		UPDATE tasks 
		SET title = $1, description = $2, status = $3, updated_at = $4
		WHERE id = $5
		RETURNING id, title, description, status, created_at, updated_at
	`

	err := r.db.QueryRow(
		query,
		task.Title,
		task.Description,
		task.Status,
		task.UpdatedAt,
		task.ID,
	).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return domain.Task{}, domain.ErrTaskNotFound
	}

	return task, err
}

func (r *TaskRepository) DeleteTask(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	query := `DELETE FROM tasks WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return domain.ErrTaskNotFound
	}

	return nil
}

func (r *TaskRepository) GetAllTasks() []domain.Task {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var tasks []domain.Task
	query := `SELECT id, title, description, status, created_at, updated_at FROM tasks`

	rows, err := r.db.Query(query)
	if err != nil {
		return tasks
	}
	defer rows.Close()

	for rows.Next() {
		var task domain.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			continue
		}
		tasks = append(tasks, task)
	}

	return tasks
}
