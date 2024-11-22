package project

import (
	"git-project-management/internal/task"
	"time"
)

type ProjectEntity struct {
	tableName   struct{}  `pg:"project"`
	ID          int64     `pg:"id,pk"`                    // Unique identifier
	Name        string    `pg:"name,notnull"`             // Project name
	Description string    `pg:"description"`              // Optional project description
	StartDate   time.Time `pg:"start_date"`               // Project start date
	EndDate     time.Time `pg:"end_date"`                 // Project end date
	CreatedBy   int64     `pg:"created_by"`               // User ID who created the project
	CreatedAt   time.Time `pg:"created_at,default:now()"` // Timestamp when the project was created
}

// ---

type ProjectDTO struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedBy   int64     `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	// task count
	// done task count
	// activity count
}
type GetAllRequest struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

type GetAllResponse struct {
	Body []ProjectDTO
}

// ---
type GetOneRequest struct {
	Id int64 `path:"id"`
}

type GetOneResponse struct {
	Body ProjectDTO
}

// ---

type GetAllTasksRequest struct {
	Limit     int   `query:"limit"`
	Offset    int   `query:"offset"`
	ProjectId int64 `path:"project_id"`
}
type GetAllTasksResponse struct {
	Body []task.TaskDTO
}
