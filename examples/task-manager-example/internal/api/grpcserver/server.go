package grpcserver

import (
	"context"
	"task-manager-example/internal/api/proto"
	"task-manager-example/internal/domain"
	"task-manager-example/internal/service"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	proto.UnimplementedTaskServiceServer
	service *service.TaskService
}

func NewServer(service *service.TaskService) *Server {
	return &Server{
		service: service,
	}
}

func (s *Server) CreateTask(ctx context.Context, req *proto.CreateTaskRequest) (*proto.CreateTaskResponse, error) {
	task := domain.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	createdTask, err := s.service.CreateTask(task)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateTaskResponse{
		Task: domainToProtoTask(createdTask),
	}, nil
}

func (s *Server) GetTask(ctx context.Context, req *proto.GetTaskRequest) (*proto.GetTaskResponse, error) {
	task, err := s.service.GetTask(req.Id)
	if err != nil {
		if err == domain.ErrTaskNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetTaskResponse{
		Task: domainToProtoTask(task),
	}, nil
}

func (s *Server) UpdateTask(ctx context.Context, req *proto.UpdateTaskRequest) (*proto.UpdateTaskResponse, error) {
	task := protoToDomainTask(req.Task)
	updatedTask, err := s.service.UpdateTask(task)
	if err != nil {
		if err == domain.ErrTaskNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.UpdateTaskResponse{
		Task: domainToProtoTask(updatedTask),
	}, nil
}

func (s *Server) DeleteTask(ctx context.Context, req *proto.DeleteTaskRequest) (*proto.DeleteTaskResponse, error) {
	err := s.service.DeleteTask(req.Id)
	if err != nil {
		if err == domain.ErrTaskNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.DeleteTaskResponse{}, nil
}

func (s *Server) GetAllTasks(ctx context.Context, req *proto.GetAllTasksRequest) (*proto.GetAllTasksResponse, error) {
	tasks := s.service.GetAllTasks()
	protoTasks := make([]*proto.Task, len(tasks))
	for i, task := range tasks {
		protoTasks[i] = domainToProtoTask(task)
	}

	return &proto.GetAllTasksResponse{
		Tasks: protoTasks,
	}, nil
}

// Helper functions to convert between domain and proto types
func domainToProtoTask(task domain.Task) *proto.Task {
	return &proto.Task{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   task.UpdatedAt.Format(time.RFC3339),
	}
}

func protoToDomainTask(task *proto.Task) domain.Task {
	createdAt, _ := time.Parse(time.RFC3339, task.CreatedAt)
	updatedAt, _ := time.Parse(time.RFC3339, task.UpdatedAt)

	return domain.Task{
		ID:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
