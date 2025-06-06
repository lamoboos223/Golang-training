# List all services
grpcurl -plaintext localhost:9090 list

# Create a new task
grpcurl -plaintext -d '{
  "title": "Learn Go",
  "description": "Study Go programming language",
  "status": false
}' localhost:9090 task.TaskService/CreateTask

# Get a task by ID
grpcurl -plaintext -d '{
  "id": "e05df603-24d6-4abf-ba48-a8fefa9ed72f"
}' localhost:9090 task.TaskService/GetTask

# Update a task
grpcurl -plaintext -d '{
  "task": {
    "id": "9313c909-b672-485c-804b-9302c08fee32",
    "title": "Learn Go - Updated",
    "description": "Study Go programming language and its concurrency features",
    "status": true
  }
}' localhost:9090 task.TaskService/UpdateTask

# Delete a task
grpcurl -plaintext -d '{
  "id": "9313c909-b672-485c-804b-9302c08fee32"
}' localhost:9090 task.TaskService/DeleteTask

# Get all tasks
grpcurl -plaintext localhost:9090 task.TaskService/GetAllTasks
