package service

import (
	"context"
	"github.com/Hellsker/Todo-List/internal/entity"
	"github.com/Hellsker/Todo-List/internal/repository"
	"time"
)

type taskService struct {
	repo repository.TaskInterface
}

func New(repo repository.TaskInterface) *taskService {
	return &taskService{repo: repo}
}

func (s *taskService) GetById(ctx context.Context, id uint64) (entity.TaskResponse, error) {
	task, err := s.repo.GetById(ctx, id)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (s *taskService) GetByDateWithFilter(ctx context.Context, date time.Time, status string) ([]entity.TaskResponse, error) {
	tasks, err := s.repo.GetByDateWithFilter(ctx, date, status)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
func (s *taskService) GetAllWithPagination(ctx context.Context, limit uint64, offset uint64, status string) ([]entity.TaskResponse, error) {
	tasks, err := s.repo.GetAllWithPagination(ctx, limit, offset, status)
	if err != nil {
		return nil, err
	}
	return tasks, nil

}
func (s *taskService) Save(ctx context.Context, task entity.TaskRequest) error {
	err := s.repo.Create(ctx, task)
	if err != nil {
		return err
	}
	return nil
}
func (s *taskService) Delete(ctx context.Context, id uint64) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *taskService) Update(ctx context.Context, id uint64, task entity.TaskRequest) error {
	err := s.repo.Update(ctx, id, task)
	if err != nil {
		return err
	}
	return nil
}
