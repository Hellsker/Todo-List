package repository

import (
	"context"
	"fmt"
	"github.com/Hellsker/Todo-List/internal/entity"
	"github.com/Hellsker/Todo-List/pkg/postgres"
	"time"
)

type TaskRepository struct {
	db *postgres.Postgres
}

// New method TaskRepository constructor
func New(db *postgres.Postgres) *TaskRepository {
	return &TaskRepository{db: db}
}

// Create method will insert new task to database. 'C' part of the CRUDL
func (r *TaskRepository) Create(ctx context.Context, task entity.TaskRequest) error {
	insertExec := "INSERT INTO tasks (title,description,due_date,completed) VALUES ($1,$2,$3,$4)"
	_, err := r.db.GetPool().Exec(ctx, insertExec, task.Title, task.Description, task.DueDate, task.Completed)
	if err != nil {
		return fmt.Errorf("TaskRepository - Create - r.db.GetPoll().Exec: %w", err)
	}
	return nil
}

// GetById method will read task to database. 'R' part of the CRUDL
func (r *TaskRepository) GetById(ctx context.Context, id uint64) (entity.TaskResponse, error) {
	selectQueryRow := "SELECT id,title,description,due_date,completed FROM tasks WHERE id = $1"
	var task entity.TaskResponse
	err := r.db.GetPool().QueryRow(ctx, selectQueryRow, id).Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed)
	if err != nil {
		return entity.TaskResponse{}, err
	}
	return task, nil
}

// GetById method will read task to database. 'R' part of the CRUDL
func (r *TaskRepository) GetByDateWithFilter(ctx context.Context, date time.Time, status string) ([]entity.TaskResponse, error) {
	fmt.Println(date)
	selectQuery := "SELECT id, title, description, due_date, completed FROM tasks WHERE due_date = $1"
	if status == "" {
		rows, err := r.db.GetPool().Query(ctx, selectQuery, date)
		if err != nil {
			return nil, fmt.Errorf("TaskRepository - GetByDateWithFilter - r.db.GetPoll().Query: %w", err)
		}
		defer rows.Close()
		var tasks []entity.TaskResponse
		for rows.Next() {
			var task entity.TaskResponse
			err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed)
			if err != nil {
				return nil, fmt.Errorf("TaskRepository - GetByDateWithFilter - rows.Next: %w", err)
			}
			tasks = append(tasks, task)
		}
		return tasks, nil
	} else {
		selectQuery += " AND completed = $2"
		rows, err := r.db.GetPool().Query(ctx, selectQuery, date, status)
		if err != nil {
			return nil, fmt.Errorf("TaskRepository -  GetByDateWithFilter - r.db.GetPoll().Query: %w", err)
		}
		defer rows.Close()
		var tasks []entity.TaskResponse
		for rows.Next() {
			var task entity.TaskResponse
			err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed)
			if err != nil {
				return nil, fmt.Errorf("TaskRepository -  GetByDateWithFilter - rows.Next: %w", err)
			}
			tasks = append(tasks, task)
		}
		return tasks, nil
	}
}

// Update method will update task to database. 'U' part of the CRUDL
func (r *TaskRepository) Update(ctx context.Context, id uint64, task entity.TaskRequest) error {
	updateExec := "UPDATE tasks SET title = $1, description = $2, due_date = $3, completed = $4 WHERE id = $5"
	_, err := r.db.GetPool().Exec(ctx, updateExec, task.Title, task.Description, task.DueDate, task.Completed, id)
	if err != nil {
		return fmt.Errorf("TaskRepository - Update - r.db.GetPoll().Exec: %w", err)
	}
	return nil
}

// Delete method will delete task to database. 'D' part of the CRUDL
func (r *TaskRepository) Delete(ctx context.Context, id uint64) error {
	deleteExec := "DELETE FROM tasks WHERE id = $1"
	_, err := r.db.GetPool().Exec(ctx, deleteExec, id)
	if err != nil {
		return fmt.Errorf("TaskRepository - Delete - r.db.GetPoll().Exec: %w", err)
	}
	return nil
}

// GetAllWithPagination method will return all tasks in the database. 'L' part of CRUDL
func (r *TaskRepository) GetAllWithPagination(ctx context.Context, limit uint64, offset uint64, status string) ([]entity.TaskResponse, error) {
	selectQuery := "SELECT id, title, description, due_date, completed FROM tasks"
	if status == "" {
		selectQuery += " LIMIT $1 OFFSET $2"
		rows, err := r.db.GetPool().Query(ctx, selectQuery, limit, offset)
		if err != nil {
			return nil, fmt.Errorf("TaskRepository - GetAllWithPagination - r.db.GetPoll().Query: %w", err)
		}
		defer rows.Close()
		var tasks []entity.TaskResponse
		for rows.Next() {
			var task entity.TaskResponse
			err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed)
			if err != nil {
				return nil, fmt.Errorf("TaskRepository - GetAllWithPagination - rows.Next: %w", err)
			}
			tasks = append(tasks, task)
		}
		return tasks, nil
	} else {
		selectQuery += " WHERE completed = $1 LIMIT $2 OFFSET $3"
		rows, err := r.db.GetPool().Query(ctx, selectQuery, status, limit, offset)
		if err != nil {
			return nil, fmt.Errorf("TaskRepository - GetAllWithPagination - r.db.GetPoll().Query: %w", err)
		}
		defer rows.Close()
		var tasks []entity.TaskResponse
		for rows.Next() {
			var task entity.TaskResponse
			err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed)
			if err != nil {
				return nil, fmt.Errorf("TaskRepository - GetAllWithPagination - rows.Next: %w", err)
			}
			tasks = append(tasks, task)
		}
		return tasks, nil
	}

}
