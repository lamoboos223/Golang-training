package domain

import "errors"

var (
	ErrTaskNotFound      = errors.New("task not found")
	ErrTaskIDRequired    = errors.New("task id is required")
	ErrTaskTitleRequired = errors.New("task title is required")
)
