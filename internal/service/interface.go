package service

import (
	"context"
	"github.com/Hellsker/Todo-List/internal/entity"
	"time"
)

// taskService interface
type TaskInterface interface {
	GetById(ctx context.Context, id uint64) (entity.TaskResponse, error)
	GetByDateWithFilter(ctx context.Context, date time.Time, status string) ([]entity.TaskResponse, error)
	GetAllWithPagination(ctx context.Context, limit uint64, offset uint64, status string) ([]entity.TaskResponse, error)
	Save(ctx context.Context, task entity.TaskRequest) error
	Delete(ctx context.Context, id uint64) error
	Update(ctx context.Context, id uint64, task entity.TaskRequest) error
}
