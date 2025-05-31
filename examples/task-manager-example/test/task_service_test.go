package test

import (
	"task-manager-example/internal/domain"
	"task-manager-example/internal/repo/memory"
	"task-manager-example/internal/service"
	"testing"
)

func setupTestService() *service.TaskService {
	repo := memory.NewTaskRepository()
	return service.NewTaskService(repo)
}

func TestTaskService_CRUD(t *testing.T) {
	service := setupTestService()

	// Test Create
	task := domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      false,
	}

	createdTask, err := service.CreateTask(task)
	if err != nil {
		t.Fatalf("Failed to create task: %v", err)
	}

	if createdTask.ID == "" {
		t.Error("Created task should have an ID")
	}
	if createdTask.Title != task.Title {
		t.Errorf("Expected title %s, got %s", task.Title, createdTask.Title)
	}

	// Test Read
	retrievedTask, err := service.GetTask(createdTask.ID)
	if err != nil {
		t.Fatalf("Failed to get task: %v", err)
	}

	if retrievedTask.ID != createdTask.ID {
		t.Errorf("Expected ID %s, got %s", createdTask.ID, retrievedTask.ID)
	}

	// Test Update
	retrievedTask.Title = "Updated Title"
	retrievedTask.Description = "Updated Description"

	updatedTask, err := service.UpdateTask(retrievedTask)
	if err != nil {
		t.Fatalf("Failed to update task: %v", err)
	}

	if updatedTask.Title != "Updated Title" {
		t.Errorf("Expected title 'Updated Title', got %s", updatedTask.Title)
	}

	// Test Delete
	err = service.DeleteTask(updatedTask.ID)
	if err != nil {
		t.Fatalf("Failed to delete task: %v", err)
	}

	_, err = service.GetTask(updatedTask.ID)
	if err != domain.ErrTaskNotFound {
		t.Error("Task should be deleted")
	}
}

func TestTaskService_GetAllTasks(t *testing.T) {
	service := setupTestService()

	// Create multiple tasks
	tasks := []domain.Task{
		{Title: "Task 1", Description: "Description 1"},
		{Title: "Task 2", Description: "Description 2"},
		{Title: "Task 3", Description: "Description 3"},
	}

	for _, task := range tasks {
		_, err := service.CreateTask(task)
		if err != nil {
			t.Fatalf("Failed to create task: %v", err)
		}
	}

	// Test GetAllTasks
	allTasks := service.GetAllTasks()
	if len(allTasks) != len(tasks) {
		t.Errorf("Expected %d tasks, got %d", len(tasks), len(allTasks))
	}

	// Verify task contents
	for i, task := range allTasks {
		if task.Title != tasks[i].Title {
			t.Errorf("Expected title %s, got %s", tasks[i].Title, task.Title)
		}
		if task.Description != tasks[i].Description {
			t.Errorf("Expected description %s, got %s", tasks[i].Description, task.Description)
		}
	}
}

func TestTaskService_ErrorCases(t *testing.T) {
	service := setupTestService()

	// Test GetTask with non-existent ID
	_, err := service.GetTask("non-existent-id")
	if err != domain.ErrTaskNotFound {
		t.Errorf("Expected ErrTaskNotFound, got %v", err)
	}

	// Test UpdateTask with non-existent ID
	nonExistentTask := domain.Task{
		ID:          "non-existent-id",
		Title:       "Test Task",
		Description: "Test Description",
	}
	_, err = service.UpdateTask(nonExistentTask)
	if err != domain.ErrTaskNotFound {
		t.Errorf("Expected ErrTaskNotFound, got %v", err)
	}

	// Test DeleteTask with non-existent ID
	err = service.DeleteTask("non-existent-id")
	if err != domain.ErrTaskNotFound {
		t.Errorf("Expected ErrTaskNotFound, got %v", err)
	}
}
