package entity

import (
	"encoding/json"
	"fmt"
	"time"
)

type TaskResponse struct {
	ID          int       `json:"id" example:"22"`
	Title       string    `json:"title" example:"Read a book"`
	Description string    `json:"description" example:"Read the book The Lord of the Rings, stopped at the third chapter"`
	DueDate     time.Time `json:"due_date" example:"2024-06-10T00:00:00Z"`
	Completed   bool      `json:"completed" example:"true"`
}

type TaskRequest struct {
	Title       string    `json:"title" example:"Read a book"`
	Description string    `json:"description" example:"Read the book The Lord of the Rings, stopped at the third chapter"`
	DueDate     time.Time `json:"due_date" example:"2024-06-10"`
	Completed   bool      `json:"completed" example:"true"`
}

func (tq *TaskRequest) UnmarshalJSON(b []byte) (err error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	tq.Title = m["title"].(string)
	tq.Description = m["description"].(string)
	dueDateString, ok := m["due_date"].(string)
	if !ok {
		return fmt.Errorf("expected due_date to be a string")
	}
	tq.DueDate, err = time.Parse(time.DateOnly, dueDateString)
	if err != nil {
		return err
	}
	tq.Completed, ok = m["completed"].(bool)
	if !ok {
		return fmt.Errorf("expected completed to be a boolean")
	}
	return nil
}
